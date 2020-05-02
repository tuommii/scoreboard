package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"miikka.xyz/scoreboard/game"
)

const (
	maxGames = 10000
)

// Server ...
type Server struct {
	// This gets passed to Game for creating ID
	counter int
	HTTP    *http.Server
	games   map[string]*game.Course
	courses []game.CourseInfo
	mu      sync.Mutex
}

// New creates a server
func New(path string) *Server {
	server := &Server{counter: 1}
	router := mux.NewRouter()
	server.HTTP = &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server.games = make(map[string]*game.Course)

	file, err := ioutil.ReadFile(path + "courses.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(file), &server.courses)
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/games/create", server.CreateGameHandle).Methods("POST")
	router.HandleFunc("/games/edit", server.EditGameHandle).Methods("POST")
	router.HandleFunc("/games/{id}", server.GetGameHandle).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(path + "public")))

	return server
}

// GetGameHandle returns game by id
func (s *Server) GetGameHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, exist := s.games[id]; !exist {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(s.games[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(bytes))
}

// CreateGameHandle creates new game
func (s *Server) CreateGameHandle(w http.ResponseWriter, r *http.Request) {
	if len(s.games) > maxGames {
		http.Error(w, "Server if full", http.StatusTooManyRequests)
		return
	}

	course, err := game.CreateFromRequest(r.Body, s.courses, s.counter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.updateCounter()

	courseJSON, err := json.Marshal(course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.games[course.ID] = course
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(courseJSON))
}

// EditGameHandle updates game on server
func (s *Server) EditGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var c *game.Course
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := c.ID
	if _, found := s.games[id]; !found {
		http.Error(w, "Game not found", http.StatusInternalServerError)
		return
	}

	if s.games[id].Active > s.games[id].BasketCount || s.games[id].Active < 1 {
		fmt.Fprintf(w, string(bytes))
		return
	}

	// If editedAt is fraud, we can still delete game with orginal createdAt
	temp := s.games[id].CreatedAt
	s.games[id] = c
	s.games[id].CreatedAt = temp

	resp, err := json.Marshal(s.games[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(resp))
}

func (s *Server) updateCounter() {
	if s.counter > maxGames {
		s.counter = 0
	}
	s.mu.Lock()
	s.counter++
	s.mu.Unlock()
}

// CleanGames removes old games
func (s *Server) CleanGames() {
	for {
		time.Sleep(20 * time.Minute)
		s.remove()
	}
}

func (s *Server) remove() {
	for id, game := range s.games {
		if time.Since(game.EditedAt) > time.Hour*1 || time.Since(game.CreatedAt) > (time.Hour*5) {
			log.Println(id, "deleted")
			s.mu.Lock()
			delete(s.games, id)
			s.mu.Unlock()
		}
	}
}
