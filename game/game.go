package game

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

// Course ...
type Course struct {
	ID          string `json:"id"`
	BasketCount int    `json:"basketCount"`
	Active      int    `json:"active"`
	// OrderNumber is the key
	Baskets   map[int]*Basket `json:"baskets"`
	CreatedAt time.Time       `json:"createdAt"`
	EditedAt  time.Time       `json:"editedAt"`
	Name      string          `json:"name"`
}

// Basket ...
type Basket struct {
	// Lets save ordernumber also here just in case
	OrderNum int `json:"orderNum"`
	Par      int `json:"par"`
	// Key is player name
	Scores map[string]*BasketScore `json:"scores"`
}

// BasketScore ...
type BasketScore struct {
	Score int `json:"score"`
	// For graph / stats ?
	Total int `json:"total"`
	OB    int `json:"ob"`
}

// NewCourse returns new Course
func NewCourse() *Course {
	// TODO: Check errors
	c := &Course{}
	baskets := make(map[int]*Basket)
	c.Baskets = baskets
	c.Active = 1
	c.CreatedAt = time.Now()
	c.EditedAt = time.Now()
	c.Name = "Default"
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
	var id string
	for _, player := range players {
		id += strings.ToLower((string(player[0])))
	}
	id += strconv.Itoa(counter)
	return id
}

// CreateCourse creates course and sets all pars to 3
func CreateCourse(players []string, baskets int, counter int) *Course {
	// TODO: check bad input
	course := NewCourse()
	course.ID = createID(players, counter)
	course.BasketCount = baskets
	for i := 1; i <= baskets; i++ {
		basket := NewBasket()
		basket.Par = 3
		basket.OrderNum = i
		for _, player := range players {
			basketScore := NewBasketScore()
			basketScore.Score = basket.Par
			// basketScore.Total = basket.Par
			basket.Scores[player] = basketScore
		}
		course.Baskets[i] = basket
	}
	return course
}

// CreateExistingCourse take's pars from real course
func CreateExistingCourse(players []string, baskets int, counter int, pars []int, name string) *Course {
	course := NewCourse()
	course.ID = createID(players, counter)
	course.Name = name
	course.BasketCount = baskets
	for i := 1; i <= baskets; i++ {
		basket := NewBasket()
		basket.Par = pars[i-1]
		basket.OrderNum = i
		for _, player := range players {
			basketScore := NewBasketScore()
			basketScore.Score = basket.Par
			// basketScore.Total = basket.Par
			basket.Scores[player] = basketScore
		}
		course.Baskets[i] = basket
	}
	return course
}
