package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"miikka.xyz/scoreboard/game"
)

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
	log.Println("created: ", course.Name, "\ngames total:", len(s.games))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(courseJSON))
}

// EditGameHandle updates game on server
func (s *Server) EditGameHandle(w http.ResponseWriter, r *http.Request) {
	c, orginal, err := game.CourseFromJSON(r.Body)
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
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(orginal))
		return
	}

	// If editedAt is fraud, we can still delete game with createdAt
	temp := s.games[id].CreatedAt
	s.games[id] = c
	s.games[id].CreatedAt = temp

	resp, err := json.Marshal(s.games[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(resp))
}

// ExitGameHandle sets HasBooker to false
func (s *Server) ExitGameHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, exist := s.games[id]; !exist {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	log.Println("ADASDSADSDSDSA")
	s.games[id].HasBooker = false
	fmt.Fprintf(w, "{}")
}
