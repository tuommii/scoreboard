package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"miikka.xyz/scoreboard/game"
	"miikka.xyz/scoreboard/geo"
)

const (
	maxBaskets   = 36
	maxPlayers   = 5
	maxPlayerLen = 10
)

// New ...
func New(path string) *Server {
	server := &Server{}
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
	w.Header().Set("Content-Type", "application/json")
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
	fmt.Fprintf(w, string(bytes))
}

// CreateGameHandle creates new game
func (s *Server) CreateGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(s.games) > 10000 {
		http.Error(w, "Server if full", http.StatusTooManyRequests)
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var query StartingRequest
	err = json.Unmarshal(bytes, &query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(query.Players) > maxPlayers || query.BasketCount > maxBaskets {
		http.Error(w, "Ivalid data", http.StatusInternalServerError)
		return
	}

	for _, player := range query.Players {
		if len(player) > maxPlayerLen {
			http.Error(w, "Ivalid data", http.StatusInternalServerError)
			return
		}
	}

	s.mu.Lock()
	s.counter++
	s.mu.Unlock()

	var course *game.Course
	for _, info := range s.courses {
		m := geo.Distance(query.Lat, query.Lon, info.Lat, info.Lon)
		if m < 1000 && m > 0 {
			course = game.CreateExistingCourse(query.Players, query.BasketCount, s.counter, info.Pars, info.ShortName)
			fmt.Println("created", info.ShortName)
			break
		}
	}

	if course == nil {
		fmt.Println("creating default (all par 3)")
		course = game.CreateCourse(query.Players, query.BasketCount, s.counter)
	}

	bytes, err = json.Marshal(course)
	var c *game.Course
	json.Unmarshal(bytes, &c)
	if err != nil {
		fmt.Fprintf(w, "{}")
		return
	}
	s.games[course.ID] = course
	fmt.Fprintf(w, string(bytes))
}

// EditGameHandle updates game on server also
func (s *Server) EditGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Read body
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Make course from body
	var c *game.Course
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if game is found
	id := c.ID
	if _, found := s.games[id]; !found {
		http.Error(w, "Game not found", http.StatusInternalServerError)
		return
	}

	// TODO: Check ip or?
	// Posted by someone else
	// if c.CreatedAt != s.games[id].CreatedAt {
	// 	http.Error(w, "Hmm...", http.StatusInternalServerError)
	// 	return
	// }

	// Not going over last basket
	if s.games[id].Active > s.games[id].BasketCount {
		fmt.Fprintf(w, string(bytes))
		return
	}

	s.mu.Lock()
	temp := s.games[id].CreatedAt
	s.games[id] = c
	s.games[id].CreatedAt = temp
	log.Println(s.games[id].EditedAt.Sub(s.games[id].CreatedAt))
	s.mu.Unlock()

	resp, err := json.Marshal(s.games[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(resp))
}

// CleanGames ...
func (s *Server) CleanGames() {
	for {
		time.Sleep(20 * time.Minute)
		s.remove()
	}
}

func (s *Server) remove() {
	for id, game := range s.games {
		if time.Since(game.EditedAt) > time.Hour*1 {
			log.Println(id, "deleted")
			s.mu.Lock()
			delete(s.games, id)
			s.mu.Unlock()
		} else if time.Since(game.CreatedAt) > (time.Hour * 5) {
			log.Println(id, "deleted")
			s.mu.Lock()
			delete(s.games, id)
			s.mu.Unlock()
		}
	}
}
