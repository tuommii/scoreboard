package game

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

var counter int

// Course ...
type Course struct {
	ID          string `json:"id"`
	BasketCount int    `json:"basketCount"`
	Active      int    `json:"active"`
	// OrderNumber is the key
	Baskets   map[int]*Basket `json:"baskets"`
	CreatedAt time.Time       `json:"-"`
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
	// For graph
	Total int `json:"total"`
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

// CreateID ...
func createID(players []string, counter int) string {
	sort.Strings(players)
	id := strconv.Itoa(counter)
	for _, player := range players {
		id += strings.ToLower((string(player[0])))
	}
	return id
}

// CreateCourse ...
func CreateCourse(players []string, baskets int, counter int) *Course {
	// TODO: check bad input
	course := NewCourse()
	course.CreatedAt = time.Now()
	course.ID = createID(players, counter)
	course.BasketCount = baskets
	course.Active = 1
	for i := 0; i < baskets; i++ {
		basket := NewBasket()
		// TODO: Just for testing
		basket.Par = 3 + i
		basket.OrderNum = i + 1
		for _, player := range players {
			basketScore := NewBasketScore()
			basketScore.Score = basket.Par
			// basketScore.Total = basket.Par
			basket.Scores[player] = basketScore
		}
		course.Baskets[i+1] = basket
	}
	return course
}
