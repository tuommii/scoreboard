package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	cmap "github.com/orcaman/concurrent-map"
	"miikka.xyz/scoreboard/game"
)

const (
	maxGames    = 10000
	maxBodySize = 1048576
)

// Server ...
type Server struct {
	HTTP *http.Server
	mu   sync.Mutex
	// counter gets passed to game for creating unique ID
	counter int
	// User created courses
	games  map[string]*game.Course
	games2 cmap.ConcurrentMap
	// Existing courses, if user is near a course, create that
	courses []game.CourseInfo
}

// New creates a server
func New(path string) *Server {
	server := &Server{counter: 1}
	server.games2 = cmap.New()
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

// SaveMemory saves data in memory to json-file
// TODO: Make this database
// TODO: Take cear of hasBooker
func (s *Server) SaveMemory(path string) {
	file, err := json.Marshal(s.games)
	if err != nil {
		log.Println("Marshaling memory data failed!", err)
		return
	}

	err = ioutil.WriteFile(path+"assets/memory.json", file, 0644)
	if err != nil {
		log.Println("Saving memory failed!", err)
	}
	log.Println(s.games2.Count(), "games saved")
}

// LoadMemory ...
// TODO: Take cear of counter value
func (s *Server) LoadMemory(path string) {
	log.Println("Loading 0 games from previous memory")
	// file, err := ioutil.ReadFile(path + "assets/memory.json")
	// if err != nil {
	// 	log.Println("Error while opening file", err)
	// 	return
	// }

	// err = json.Unmarshal(file, &s.games)
	// if err != nil {
	// 	log.Println("Error while unmarshaling previous memory", err)
	// 	return
	// }

	// log.Println("loaded ", len(s.games))
}
