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
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	s.rw.RLock()
	if _, exist := s.games[id]; !exist {
		http.Error(w, jsonErr("Not found"), http.StatusInternalServerError)
		s.rw.RUnlock()
		return
	}

	bytes, err := json.Marshal(s.games[id])
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		s.rw.RUnlock()
		return
	}
	s.rw.RUnlock()
	fmt.Fprintf(w, string(bytes))
}

// CreateGameHandle creates new game
func (s *Server) CreateGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s.rw.RLock()
	if len(s.games) > maxGames {
		http.Error(w, jsonErr("Server is full"), http.StatusTooManyRequests)
		s.rw.RUnlock()
		return
	}

	course, err := game.CreateFromRequest(r.Body, s.courses, s.counter)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		s.rw.RUnlock()
		return
	}
	s.rw.RUnlock()

	s.rw.Lock()
	s.updateCounter()
	s.rw.Unlock()

	courseJSON, err := json.Marshal(course)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	s.rw.RLock()
	log.Println("created: ", course.ID, course.Name, "\ngames total:", len(s.games))
	s.rw.RUnlock()

	s.rw.Lock()
	s.games[course.ID] = course
	s.rw.Unlock()

	fmt.Fprintf(w, string(courseJSON))
}

// EditGameHandle updates game on server
func (s *Server) EditGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c, orginal, err := game.CourseFromJSON(r.Body)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}

	id := c.ID
	s.rw.RLock()
	if _, found := s.games[id]; !found {
		http.Error(w, jsonErr("Game not found"), http.StatusInternalServerError)
		s.rw.RUnlock()
		return
	}

	if s.games[id].Active > s.games[id].BasketCount || s.games[id].Active < 1 {
		fmt.Fprintf(w, string(orginal))
		s.rw.RUnlock()
		return
	}

	s.rw.RUnlock()
	s.rw.Lock()
	// If editedAt is fraud, we can still delete game with createdAt
	temp := s.games[id].CreatedAt
	s.games[id] = c
	s.games[id].CreatedAt = temp
	s.rw.Unlock()

	s.rw.RLock()
	resp, err := json.Marshal(s.games[id])
	if err != nil {
		s.rw.RUnlock()
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	s.rw.RUnlock()
	fmt.Fprintf(w, string(resp))
}

// ExitGameHandle sets HasBooker to false
func (s *Server) ExitGameHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	s.rw.RLock()
	if _, exist := s.games[id]; !exist {
		http.Error(w, jsonErr("Not found"), http.StatusInternalServerError)
		s.rw.RUnlock()
		return
	}
	s.rw.RUnlock()

	s.rw.Lock()
	s.games[id].HasBooker = false
	s.rw.Unlock()
	fmt.Fprintf(w, "{}")
}
