package manager

import (
	"encoding/json"
	"io"
	"log"
	"strconv"
	"strings"

	"miikka.xyz/sgoreboard/game"
)

// CreateCourse ...
func CreateCourse(players []string, baskets int) *game.Course {
	// TODO: check bad input
	course := game.NewCourse()
	for i := 0; i < baskets; i++ {
		basket := game.NewBasket()
		basket.OrderNum = i + 1
		for _, player := range players {
			basketScore := game.NewBasketScore()
			basket.Scores[player] = basketScore
		}
		course.Baskets[i+1] = basket
	}
	return course
}

// JSONToCourse ...
func JSONToCourse(data string) *game.Course {
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

	course := game.NewCourse()
	for jsonKey, jsonValue := range result {
		basket := game.NewBasket()

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
					bs := &game.BasketScore{}
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
