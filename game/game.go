package game

import (
	"encoding/json"
	"io"
	"log"
	"strconv"
	"strings"
)

var counter int

// Course ...
type Course struct {
	ID          string `json:"id"`
	BasketCount int    `json:"basketCount"`
	// OrderNumber is the key
	Baskets map[int]*Basket
}

// Basket ...
type Basket struct {
	// Lets save ordernumber also here
	OrderNum int
	Par      int
	// Key is player name
	Scores map[string]*BasketScore
}

// BasketScore ...
type BasketScore struct {
	Score int
	OB    int
}

// NewCourse returns new *Course
func NewCourse() *Course {
	// TODO: Check errors
	counter++
	c := &Course{ID: strconv.Itoa(counter)}
	baskets := make(map[int]*Basket)
	c.Baskets = baskets
	return c
}

func NewBasket() *Basket {
	basket := &Basket{}
	scores := make(map[string]*BasketScore)
	basket.Scores = scores
	return basket
}

func NewBasketScore() *BasketScore {
	basketScore := &BasketScore{}
	return basketScore
}

func JsonToCourse(data string) *Course {
	var result map[string]interface{}

	dec := json.NewDecoder(strings.NewReader(data))
	for {
		// var ls LaneScores
		err := dec.Decode(&result)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
	}

	course := NewCourse()
	for jsonKey, jsonValue := range result {
		basket := NewBasket()

		if jsonKey == "basketCount" {
			course.BasketCount = int(jsonValue.(float64))
		} else if jsonKey == "id" {
			course.ID = jsonValue.(string)
		} else if value, err := strconv.Atoi(jsonKey); err == nil {
			// It can be parsed to number so its baskets order number
			basket.OrderNum = value

			basketJSON := result[jsonKey].(map[string]interface{})
			for basketKey, basketValue := range basketJSON {
				if basketKey == "par" {
					basket.Par = int(basketValue.(float64))
				} else {
					// fmt.Printf("UUS: %s:%+v\n", basketKey, basketValue)

					test := basketValue.(map[string]interface{})
					bs := &BasketScore{}
					for key, val := range test {
						if key == "score" {
							sc := val.(float64)
							bs.Score = int(sc)
						} else if key == "ob" {
							ob := val.(float64)
							bs.OB = int(ob)
							// bs.OB = val.(float64)
						} else {
						}
					}
					basket.Scores[basketKey] = bs
				}
			}
		} else {
			// Most outside key isn number or specified
		}
		if basket.OrderNum != 0 {
			course.Baskets[basket.OrderNum] = basket
		}
	}
	return course
}
