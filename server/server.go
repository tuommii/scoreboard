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
	server.loadCourseTemplates(path)
	router.HandleFunc("/games/create", server.CreateGameHandle).Methods("POST")
	router.HandleFunc("/games/edit", server.EditGameHandle).Methods("POST")
	router.HandleFunc("/games/{id}", server.GetGameHandle).Methods("GET")
	router.HandleFunc("/exit/{id}", server.ExitGameHandle).Methods("GET")
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

// Worker for future use
func (s *Server) Worker(lat float64, lon float64) {
	log.Println("simulating api request with:", lat, lon)
	time.Sleep(time.Second * 10)
	log.Println("api simulation done!")
}

func (s *Server) updateCounter() {
	if s.counter > maxGames {
		s.counter = 1
	}
	s.counter++
}

func (s *Server) clean(editedAgo time.Duration, createdAgo time.Duration) {
	s.rw.Lock()
	defer s.rw.Unlock()
	for id, game := range s.games {
		if time.Since(game.EditedAt) > editedAgo || time.Since(game.CreatedAt) > createdAgo {
			delete(s.games, id)
			log.Println("deleted", id, game.Name)
		}
	}
}

func (s *Server) loadCourseTemplates(path string) {
	file, err := ioutil.ReadFile(path + "assets/courses.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(file), &s.courses)
	if err != nil {
		log.Fatal(err)
	}
}

func jsonErr(msg string) string {
	return fmt.Sprintf(`{"err":"%s"}`, msg)
}
