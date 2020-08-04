package game

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"regexp"
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

// Create ...
func Create(basis Basis, designs []Design, counter int) (*Course, error) {
	if !validateBasis(basis) {
		return nil, errors.New("Invalid basis")
	}
	return create(designs, basis.Lat, basis.Lon, basis.Players, basis.BasketCount, counter), nil
}

// newCourse returns new Course
func newCourse() *Course {
	c := &Course{}
	baskets := make(map[int]*Basket)
	c.Baskets = baskets
	c.Active = 1
	c.CreatedAt = time.Now()
	c.EditedAt = c.CreatedAt
	c.Name = "Default"
	c.HasBooker = true
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

// CreateID creates unique id
func createID(players []string, counter int) string {
	sort.Strings(players)
	var id string
	for _, player := range players {
		id += strings.ToLower((string(player[0])))
	}
	id += strconv.Itoa(counter)
	return id
}

// TODO: Refactor more, too many params
// create new course
func create(designs []Design, lat float64, lon float64, players []string, basketCount int, counter int) *Course {
	for _, temp := range designs {
		m := geo.Distance(lat, lon, temp.Lat, temp.Lon)
		if m < near && m >= 0 {
			return createExistingCourse(players, counter, temp.Pars, temp.ShortName)
		}
	}
	return createCourse(players, basketCount, counter)
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

// createExistingCourse takes pars from real course
func createExistingCourse(players []string, counter int, pars []int, name string) *Course {
	basketCount := len(pars)
	c := createCourse(players, basketCount, counter)
	c.Name = name
	c.BasketCount = basketCount
	c.Name = name
	for i := 1; i <= c.BasketCount; i++ {
		c.Baskets[i].Par = pars[i-1]
		for _, player := range players {
			c.Baskets[i].Scores[player].Score = pars[i-1]
		}
	}
	return c
}

func validateBasis(basis Basis) bool {
	if len(basis.Players) > maxPlayers || basis.BasketCount > maxBaskets {
		return false
	}

	for _, player := range basis.Players {
		if len(player) > maxPlayerLen {
			return false
		}
	}
	return true
}

// LoadDesigns loads pre-made course data
func LoadDesigns(path string) []Design {
	var designs []Design
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(file), &designs)
	if err != nil {
		log.Fatal(err)
	}
	return designs
}

// AtoiID should use following re:
// re := regexp.MustCompile("[0-9]+")
func AtoiID(id string, re *regexp.Regexp) int {
	i, err := strconv.Atoi(re.FindString(id))
	if err != nil {
		i = 0
		log.Println(err, i)
	}
	return i
}
