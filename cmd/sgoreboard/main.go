package main

import (
	"io/ioutil"
	"log"
	"os"

	"miikka.xyz/sgoreboard/game"
)

func main() {
	// server.Start()
	jsonFile, _ := os.Open("./example.json")
	// paska := game.NewCourse()
	// game2 game.Course
	bytes, _ := ioutil.ReadAll(jsonFile)

	paska2 := game.JsonToCourse(string(bytes))
	log.Printf("%+v", paska2)

	// fmt.Fprintf(w, "Tiger King says Hello!")

}
