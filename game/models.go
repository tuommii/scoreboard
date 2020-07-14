package game

import "time"

// Course has many baskets
type Course struct {
	ID          string          `json:"id"`
	BasketCount int             `json:"basketCount"`
	Active      int             `json:"active"`
	Baskets     map[int]*Basket `json:"baskets"`
	CreatedAt   time.Time       `json:"createdAt"`
	EditedAt    time.Time       `json:"editedAt"`
	Name        string          `json:"name"`
	HasBooker   bool            `json:"hasBooker"`
}

// Basket has a Score struct for each player
type Basket struct {
	// Lets save ordernumber also here just in case
	OrderNum int `json:"orderNum"`
	Par      int `json:"par"`
	// Key is player name
	Scores map[string]*BasketScore `json:"scores"`
}

// BasketScore holds player stats for that basket
type BasketScore struct {
	Score int `json:"score"`
	// For graph / stats ?
	Total int `json:"total"`
	OB    int `json:"ob"`
}

// Basis holds data that is needed to create a new game
type Basis struct {
	BasketCount int      `json:"basketCount"`
	Players     []string `json:"players"`
	Lat         float64  `json:"lat"`
	Lon         float64  `json:"lon"`
}

// Design holds course related data
type Design struct {
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
