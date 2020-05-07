package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"miikka.xyz/scoreboard/game"
)

func (s *Server) getGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	s.rw.RLock()
	defer s.rw.RUnlock()
	if _, exist := s.games[id]; !exist {
		http.Error(w, jsonErr("Not found"), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(s.games[id])
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(bytes))
}

func (s *Server) createGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	start := time.Now()
	s.rw.RLock()
	defer s.rw.RUnlock()

	if len(s.games) >= maxGames {
		http.Error(w, jsonErr("Server is full"), http.StatusTooManyRequests)
		return
	}

	basis, err := parseBasis(r.Body)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	course, err := game.Create(basis, s.courses, s.counter)
	s.updateCounter()

	// Example for future use
	// go s.worker(basis.Lat, basis.Lon)

	courseJSON, err := json.Marshal(course)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	s.games[course.ID] = course
	log.Printf("%s #%s created, [%dms] len(%d)", course.Name, course.ID, time.Since(start).Milliseconds(), len(s.games))
	fmt.Fprintf(w, string(courseJSON))
}

func (s *Server) editGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c, orginal, err := gameFromJSON(r.Body)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}

	id := c.ID
	s.rw.Lock()
	defer s.rw.Unlock()
	if _, found := s.games[id]; !found {
		http.Error(w, jsonErr("Game not found"), http.StatusInternalServerError)
		return
	}

	if s.games[id].Active > s.games[id].BasketCount || s.games[id].Active < 1 {
		fmt.Fprintf(w, string(orginal))
		return
	}

	// If editedAt is a fraud, we can still delete game with createdAt
	temp := s.games[id].CreatedAt
	s.games[id] = c
	s.games[id].CreatedAt = temp

	resp, err := json.Marshal(s.games[id])
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(resp))
}

func (s *Server) exitGameHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	s.rw.Lock()
	defer s.rw.Unlock()
	if _, exist := s.games[id]; !exist {
		http.Error(w, jsonErr("Not found"), http.StatusInternalServerError)
		return
	}

	s.games[id].HasBooker = false
	fmt.Fprintf(w, "{}")
}

func (s *Server) statusHandle(w http.ResponseWriter, r *http.Request) {
	text(w, http.StatusOK, "OK")
}
