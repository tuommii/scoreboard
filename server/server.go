package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/gorilla/mux"
	cmap "github.com/orcaman/concurrent-map"
	"miikka.xyz/scoreboard/game"
)

const (
	maxGames = 10000
	// 64KB
	maxBodySize = 64_000
)

// Server ...
type Server struct {
	HTTP *http.Server
	// counter gets passed to game for creating unique ID
	counter int
	mu      sync.Mutex
	// User created courses, each item has its own mutex
	games cmap.ConcurrentMap
	// Existing courses, if user is near a course, create that
	designs []game.Design
}

// New creates a server
func New(path string) *Server {
	server := &Server{counter: 1}
	server.games = cmap.New()
	router := mux.NewRouter()
	server.HTTP = &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server.designs = game.LoadDesigns(path)
	router.HandleFunc("/games/{id}", server.getGameHandle).Methods("GET")
	router.HandleFunc("/games/{id}", server.exitGameHandle).Methods("DELETE")
	router.HandleFunc("/games", server.createGameHandle).Methods("POST")
	router.HandleFunc("/games", server.editGameHandle).Methods("PUT")
	router.HandleFunc("/_status", server.statusHandle).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(path + "public")))
	return server
}

// AutoClean removes old games from memory
func (s *Server) AutoClean(interval time.Duration, editedAlive time.Duration, createdAlive time.Duration) {
	for {
		time.Sleep(interval)
		s.clean(editedAlive, createdAlive)
	}
}

// SaveMemory saves data in memory to json-file
func (s *Server) SaveMemory(path string) {
	var arr []*game.Course

	for temp := range s.games.IterBuffered() {
		g := temp.Val.(*game.Course)
		arr = append(arr, g)
	}

	file, err := json.Marshal(&arr)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(path+"assets/memory.json", file, 0644)
	if err != nil {
		log.Println("Saving memory failed!", err)
	}
	log.Println(s.games.Count(), "games saved")
}

// LoadMemory ...
func (s *Server) LoadMemory(path string) {
	var largest int
	games := make([]*game.Course, 0)
	re := regexp.MustCompile("[0-9]+")
	file, err := ioutil.ReadFile(path + "assets/memory.json")
	if err != nil {
		log.Println("Error while opening file", err)
		return
	}

	err = json.Unmarshal(file, &games)
	if err != nil {
		log.Println("Error while unmarshaling previous memory", err)
		return
	}

	for _, g := range games {
		g.HasBooker = false
		id := game.AtoiID(g.ID, re)
		if id > largest {
			largest = id
		}
		s.games.Set(g.ID, g)
	}

	s.mu.Lock()
	s.counter = largest + 1
	if largest >= maxGames {
		s.counter = 1
	}
	s.mu.Unlock()

	log.Println("Loaded", s.games.Count(), "games from previous memory")
	log.Println("Counter is", s.counter)
}
