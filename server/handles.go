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

	// s.rw.RLock()
	// defer s.rw.RUnlock()
	temp, exist := s.games2.Get(id)
	if !exist {
		http.Error(w, jsonErr("Not found"), http.StatusInternalServerError)
		return
	}

	// tmp := temp.(*game.Course)
	bytes, err := json.Marshal(temp.(*game.Course))
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(bytes))
}

func (s *Server) createGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	start := time.Now()
	// s.rw.RLock()
	// defer s.rw.RUnlock()

	if s.games2.Count() >= maxGames {
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
	s.games2.Set(course.ID, course)
	// s.games[course.ID] = course
	log.Printf("%s #%s created, [%dms] len(%d)", course.Name, course.ID, time.Since(start).Milliseconds(), s.games2.Count())
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
	// s.rw.Lock()
	// defer s.rw.Unlock()
	// if _, found := s.games[id]; !found {
	// 	http.Error(w, jsonErr("Game not found"), http.StatusInternalServerError)
	// 	return
	// }
	tmp, exist := s.games2.Get(id)
	g := tmp.(*game.Course)
	if !exist {
		http.Error(w, jsonErr("Not found"), http.StatusInternalServerError)
		return
	}

	if g.Active > g.BasketCount || g.Active < 1 {
		fmt.Fprintf(w, string(orginal))
		return
	}

	// If editedAt is a fraud, we can still delete game with createdAt
	// temp := g.CreatedAt
	c.CreatedAt = g.CreatedAt
	s.games2.Set(id, c)
	// s.games[id] = c
	// s.games[id].CreatedAt = temp

	resp, err := json.Marshal(g)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(resp))
}

func (s *Server) exitGameHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// s.rw.Lock()
	// defer s.rw.Unlock()
	temp, exist := s.games2.Get(id)
	if !exist {
		http.Error(w, jsonErr("Not found"), http.StatusInternalServerError)
		return
	}
	g := temp.(*game.Course)
	g.HasBooker = false
	s.games2.Set(id, g)
	fmt.Fprintf(w, "{}")
}

func (s *Server) statusHandle(w http.ResponseWriter, r *http.Request) {
	text(w, http.StatusOK, "OK")
}
