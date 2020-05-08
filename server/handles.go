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

	course, exist := s.games.Get(id)
	if !exist {
		http.Error(w, jsonErr("Not found"), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(course.(*game.Course))
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(bytes))
}

func (s *Server) createGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	start := time.Now()

	if s.games.Count() >= maxGames {
		http.Error(w, jsonErr("Server is full"), http.StatusTooManyRequests)
		return
	}

	basis, err := parseBasis(r.Body)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	course, err := game.Create(basis, s.designs, s.counter)
	s.updateCounter()

	// Example for future use
	// go s.worker(basis.Lat, basis.Lon)

	courseJSON, err := json.Marshal(course)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	s.games.Set(course.ID, course)
	log.Printf("%s #%s created, [%dms] len(%d)", course.Name, course.ID, time.Since(start).Milliseconds(), s.games.Count())
	fmt.Fprintf(w, string(courseJSON))
}

func (s *Server) editGameHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c, orginal, err := gameFromJSON(r.Body)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}

	tmp, exist := s.games.Get(c.ID)
	if !exist {
		http.Error(w, jsonErr("Not found"), http.StatusInternalServerError)
		return
	}
	course := tmp.(*game.Course)

	if course.Active > course.BasketCount || course.Active < 1 {
		fmt.Fprintf(w, string(orginal))
		return
	}

	// If editedAt is a fraud, we can still delete game with createdAt
	c.CreatedAt = course.CreatedAt
	c.HasBooker = true
	s.games.Set(c.ID, c)

	resp, err := json.Marshal(c)
	if err != nil {
		http.Error(w, jsonErr(err.Error()), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(resp))
}

func (s *Server) exitGameHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	temp, exist := s.games.Get(id)
	if !exist {
		http.Error(w, jsonErr("Not found"), http.StatusInternalServerError)
		return
	}
	g := temp.(*game.Course)
	g.HasBooker = false
	s.games.Set(id, g)
	fmt.Fprintf(w, "{}")
}

func (s *Server) statusHandle(w http.ResponseWriter, r *http.Request) {
	text(w, http.StatusOK, "OK")
}
