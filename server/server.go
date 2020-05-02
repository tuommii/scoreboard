package server

import (
	"encoding/json"
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
	// counter gets passed to game for creating unique ID
	counter int
	HTTP    *http.Server
	// User created courses
	games map[string]*game.Course
	// Existing courses, if user is near a course, create that
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

	server.loadCourseTemplates(path)
	router.HandleFunc("/games/create", server.CreateGameHandle).Methods("POST")
	router.HandleFunc("/games/edit", server.EditGameHandle).Methods("POST")
	router.HandleFunc("/games/{id}", server.GetGameHandle).Methods("GET")
	router.HandleFunc("/exit/{id}", server.ExitGameHandle).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(path + "public")))

	return server
}

// AutoClean removes old games
func (s *Server) AutoClean() {
	for {
		time.Sleep(20 * time.Minute)
		s.clean()
	}
}

func (s *Server) loadCourseTemplates(path string) {
	file, err := ioutil.ReadFile(path + "courses.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(file), &s.courses)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) updateCounter() {
	s.mu.Lock()
	if s.counter > maxGames {
		s.counter = 1
	}
	s.counter++
	s.mu.Unlock()
}

func (s *Server) clean() {
	for id, game := range s.games {
		if time.Since(game.EditedAt) > time.Hour*1 || time.Since(game.CreatedAt) > (time.Hour*5) {
			log.Println("deleted", id, game.Name)
			// s.mu.Lock()
			delete(s.games, id)
			// s.mu.Unlock()
		}
	}
}
