package server

import (
	"net/http"
	"sync"

	"miikka.xyz/sgoreboard/game"
)

// Server ...
type Server struct {
	// This gets passed to Game for creating ID
	counter int
	HTTP    *http.Server
	games   map[string]*game.Course
	courses []CourseInfo
	mu      sync.Mutex
}

// StartingRequest holds data thats needed for starting new game
type StartingRequest struct {
	BasketCount int      `json:"basketCount"`
	Players     []string `json:"players"`
	Lat         float64  `json:"lat"`
	Lon         float64  `json:"lon"`
}

// CourseInfo holds course related data
type CourseInfo struct {
	ID          string  `json:"id,omitempty"`
	CountryCode string  `json:"countryCode,omitempty"`
	City        string  `json:"city,omitempty"`
	Lanes       int     `json:"lanes,omitempty"`
	Lon         float64 `json:"lon,omitempty"`
	Lat         float64 `json:"lat,omitempty"`
	ShortName   string  `json:"shortName,omitempty"`
	FullName    string  `json:"fullName,omitempty"`
	Pars        []int   `json:"pars,omitempty"`
}
