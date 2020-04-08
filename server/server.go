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
	// This gets passed to Game for creating ID
	http  *http.Server
	games map[string]*game.Course
}

// CreateQuery ...
type CreateQuery struct {
	BasketCount int      `json:"basketCount"`
	Players     []string `json:"players"`
}

// Start ...
func Start() {
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
	router.HandleFunc("/test_create", TestCreate).Methods("POST")
	router.HandleFunc("/test_edit", TestEdit).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	server.http.ListenAndServe()
	// router.HandleFunc("/games/{id:[0-9]+}", QueryGame)
}

// TestCreate ...
func TestCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var query CreateQuery
	bytes, err := ioutil.ReadAll(r.Body)
	log.Println(string(bytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bytes, &query)
	log.Printf("%+v\n%+v\n", query, err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	course := manager.CreateCourse(query.Players, query.BasketCount)

	bytes, err = json.Marshal(course)
	if err != nil {
		log.Println("DSADSADSDASSADSA")
		fmt.Fprintf(w, "{}")
	}
	fmt.Fprintf(w, string(bytes))
}

func TestEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course *game.Course
	bytes, err := ioutil.ReadAll(r.Body)
	log.Println(string(bytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bytes, &course)
	log.Printf("%+v\n", course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(bytes))
}

//
//
//
// Old code below, left for example
//
//
//

// CreateGameHandle ...
func (s *Server) CreateGameHandle(w http.ResponseWriter, r *http.Request) {
	// TODO: Mutex here
	g := game.NewCourse()
	err := json.NewDecoder(r.Body).Decode(g)
	if err != nil {
		log.Println(err)
		text(w, http.StatusBadRequest, err.Error())
		return
	}
	s.games[g.ID] = g
	fmt.Fprintf(w, "New Game: %d, %+v, %+v", len(g.Baskets), g, g.Baskets[1])
}

// GetGameHandle ...
func (s *Server) GetGameHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, exist := s.games[id]; exist {
		fmt.Fprintf(w, "%+v, %+v", s.games[id], s.games[id].Baskets)
		return
	}
	fmt.Fprintf(w, "No game found")
}

// SetBasketScore ...
func (s *Server) SetBasketScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println(vars)
	b := game.Basket{}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v", b)
}

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	text(w, 200, "OK")
}

// QueryGame ...
func QueryGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println(mux.Vars(r))
	text(w, 200, "OK")
}

func text(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}
