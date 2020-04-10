package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"miikka.xyz/sgoreboard/game"
	"miikka.xyz/sgoreboard/manager"
)

// Server ...
type Server struct {
	// ID
	counter int
	// This gets passed to Game for creating ID
	http  *http.Server
	games map[string]*game.Course
}

// StartingRequest holds data thats needed for starting new game
type StartingRequest struct {
	BasketCount int      `json:"basketCount"`
	Players     []string `json:"players"`
}

// Start ...
func Start(path string) {
	server := Server{}
	router := mux.NewRouter()
	server.http = &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Our games/courses
	server.games = make(map[string]*game.Course)

	// Init routes
	router.HandleFunc("/games/{id}/{active:[0-9]+}", server.GetGameHandle).Methods("GET")
	router.HandleFunc("/test_create", server.TestCreate).Methods("POST")
	router.HandleFunc("/test_edit", TestEdit).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(path)))
	server.http.ListenAndServe()
}

// GetGameHandle ...
func (s *Server) GetGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	active := vars["active"]

	log.Println("REQUEST", id, active)

	if _, exist := s.games[id]; exist {
		fmt.Fprintf(w, "{}")
		return
	}
	fmt.Fprintf(w, "No game found")
}

// TestCreate ...
func (s *Server) TestCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var query StartingRequest
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(bytes, &query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Inc only if all is legal
	s.counter++
	course := manager.CreateCourse(query.Players, query.BasketCount, s.counter)

	bytes, err = json.Marshal(course)
	var c *game.Course
	json.Unmarshal(bytes, &c)
	log.Println(c)
	if err != nil {
		fmt.Fprintf(w, "{}")
		return
	}
	// log.Println(string(bytes))
	s.games[course.ID] = course
	fmt.Fprintf(w, string(bytes))
}

// TestEdit ...
func TestEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("REQ:", string(bytes))

	// course := manager.JSONToCourse(string(bytes))
	var c *game.Course
	json.Unmarshal(bytes, &c)

	c.Active++
	resp, _ := json.Marshal(c)
	fmt.Printf("\n\nRESP:%+v\n", c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(resp))
}
