package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"miikka.xyz/scoreboard/game"
)

const (
	maxGames    = 10000
	maxBodySize = 1048576
)

// Server ...
type Server struct {
	HTTP *http.Server
	rw   sync.RWMutex
	// counter gets passed to game for creating unique ID
	counter int
	// User created courses
	games map[string]*game.Course
	// Existing courses, if user is near a course, create that
	courses []game.CourseInfo
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
	server.courses = game.LoadCourseTemplates(path)
	router.HandleFunc("/games/create", server.createGameHandle).Methods("POST")
	router.HandleFunc("/games/edit", server.editGameHandle).Methods("POST")
	router.HandleFunc("/games/{id}", server.getGameHandle).Methods("GET")
	router.HandleFunc("/exit/{id}", server.exitGameHandle).Methods("GET")
	router.HandleFunc("/_status", server.statusHandle).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(path + "public")))
	return server
}

// AutoClean removes old games from memory
func (s *Server) AutoClean(interval time.Duration, editedAgo time.Duration, createdAgo time.Duration) {
	for {
		time.Sleep(interval)
		s.clean(editedAgo, createdAgo)
	}
}
