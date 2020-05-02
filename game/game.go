package game

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"

	"miikka.xyz/scoreboard/geo"
)

const (
	maxBaskets   = 36
	maxPlayers   = 5
	maxPlayerLen = 10
	// max distance for existing course in meters
	near = 1000
)

// CreateFromRequest creates a new course from http request
func CreateFromRequest(body io.ReadCloser, templates []CourseInfo, counter int) (*Course, error) {
	var course *Course

	query, err := getStartingQuery(body)
	if err != nil {
		return nil, err
	}

	if !isValid(query) {
		return nil, errors.New("Invalid data")
	}

	for _, temp := range templates {
		m := geo.Distance(query.Lat, query.Lon, temp.Lat, temp.Lon)
		if m < near && m > 0 {
			course = createExistingCourse(query.Players, query.BasketCount, counter, temp.Pars, temp.ShortName)
			return course, nil
		}
	}
	return createCourse(query.Players, query.BasketCount, counter), nil
}

// CourseFromJSON creates a Course from JSON
func CourseFromJSON(body io.ReadCloser) (*Course, []byte, error) {
	var c *Course
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return c, bytes, err
	}

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return c, bytes, err
	}
	return c, bytes, nil
}

// newCourse returns new Course
func newCourse() *Course {
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

// newBasket ...
func newBasket() *Basket {
	basket := &Basket{Par: 3}
	scores := make(map[string]*BasketScore)
	basket.Scores = scores
	return basket
}

// newBasketScore ...
func newBasketScore() *BasketScore {
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
func createCourse(players []string, baskets int, counter int) *Course {
	// TODO: check bad input
	course := newCourse()
	course.ID = createID(players, counter)
	course.BasketCount = baskets
	for i := 1; i <= baskets; i++ {
		basket := newBasket()
		basket.Par = 3
		basket.OrderNum = i
		for _, player := range players {
			basketScore := newBasketScore()
			basketScore.Score = basket.Par
			// basketScore.Total = basket.Par
			basket.Scores[player] = basketScore
		}
		course.Baskets[i] = basket
	}
	return course
}

// createExistingCourse take's pars from real course
func createExistingCourse(players []string, baskets int, counter int, pars []int, name string) *Course {
	course := newCourse()
	course.ID = createID(players, counter)
	course.Name = name
	course.BasketCount = baskets
	for i := 1; i <= baskets; i++ {
		basket := newBasket()
		basket.Par = pars[i-1]
		basket.OrderNum = i
		for _, player := range players {
			basketScore := newBasketScore()
			basketScore.Score = basket.Par
			// basketScore.Total = basket.Par
			basket.Scores[player] = basketScore
		}
		course.Baskets[i] = basket
	}
	return course
}

func isValid(query CreateRequest) bool {
	if len(query.Players) > maxPlayers || query.BasketCount > maxBaskets {
		return false
	}

	for _, player := range query.Players {
		if len(player) > maxPlayerLen {
			return false
		}
	}
	return true
}

func getStartingQuery(body io.ReadCloser) (CreateRequest, error) {
	var query CreateRequest

	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return query, err
	}
	err = json.Unmarshal(bytes, &query)
	if err != nil {
		return query, err
	}
	return query, nil
}
