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
)

const (
	maxBaskets = 36
	maxPlayers = 5
)

// Server ...
type Server struct {
	// ID
	counter int
	// This gets passed to Game for creating ID
	Http  *http.Server
	games map[string]*game.Course
}

// StartingRequest holds data thats needed for starting new game
type StartingRequest struct {
	BasketCount int      `json:"basketCount"`
	Players     []string `json:"players"`
}

// New ...
func New(path string) *Server {
	server := &Server{}
	router := mux.NewRouter()
	server.Http = &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// Our games/courses
	server.games = make(map[string]*game.Course)

	router.HandleFunc("/games/{id}", server.GetGameHandle).Methods("GET")
	router.HandleFunc("/test_create", server.TestCreate).Methods("POST")
	router.HandleFunc("/test_edit", server.TestEdit).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(path)))
	return server
}

// GetGameHandle ...
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

	// Validate
	if len(query.Players) > maxPlayers && query.BasketCount > maxBaskets {
		http.Error(w, "Ivalid data", http.StatusInternalServerError)
		return
	}

	// TODO: Inc only if all is legal
	s.counter++
	course := game.CreateCourse(query.Players, query.BasketCount, s.counter)

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
func (s *Server) TestEdit(w http.ResponseWriter, r *http.Request) {
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
	if c.CreatedAt != s.games[id].CreatedAt {
		http.Error(w, "Hmm...", http.StatusInternalServerError)
		return
	}

	// Not going over last basket
	if s.games[id].Active >= s.games[id].BasketCount {
		fmt.Fprintf(w, string(bytes))
		return
	}

	// Update our internal game
	s.games[id] = c

	// s.games[id].Active++
	// // Easier read and write
	// active := s.games[id].Active

	// // Init data for next basket
	// par := s.games[id].Baskets[active].Par
	// for player := range s.games[id].Baskets[active].Scores {
	// 	// Score defaults to par
	// 	s.games[id].Baskets[active].Scores[player].Score = par
	// 	// Calc total
	// 	s.games[id].Baskets[active].Scores[player].Total += s.games[id].Baskets[active-1].Scores[player].Total
	// }

	// Edited Course to json
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
		time.Sleep(15 * time.Minute)
		s.remove()
	}
}

func (s *Server) remove() {
	for id, game := range s.games {
		if time.Since(game.CreatedAt) > (time.Hour * 6) {
			log.Println(id, "deleted")
			delete(s.games, id)
		}
	}
}
