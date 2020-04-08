package game

import (
	"strconv"
)

var counter int

// Course ...
type Course struct {
	ID          string `json:"id"`
	BasketCount int    `json:"basketCount"`
	Active      int    `json:"active"`
	// OrderNumber is the key
	Baskets map[int]*Basket `json:"baskets"`
}

// Basket ...
type Basket struct {
	// Lets save ordernumber also here
	OrderNum int `json:"orderNum"`
	Par      int `json:"par"`
	// Key is player name
	Scores map[string]*BasketScore `json:"scores"`
}

// BasketScore ...
type BasketScore struct {
	Score int `json:"score"`
	OB    int `json:"ob"`
}

// NewCourse returns new *Course
func NewCourse() *Course {
	// TODO: Check errors
	counter++
	c := &Course{ID: strconv.Itoa(counter)}
	baskets := make(map[int]*Basket)
	c.Baskets = baskets
	c.Active = 1
	return c
}

// NewBasket ...
func NewBasket() *Basket {
	basket := &Basket{Par: 3}
	scores := make(map[string]*BasketScore)
	basket.Scores = scores
	return basket
}

// NewBasketScore ...
func NewBasketScore() *BasketScore {
	basketScore := &BasketScore{}
	return basketScore
}
