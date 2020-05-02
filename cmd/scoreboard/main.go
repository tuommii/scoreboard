package main

import (
	"flag"
	"log"

	"miikka.xyz/scoreboard/server"
)

func main() {
	dir := flag.String("dir", "./", "Path to static dir")
	flag.Parse()
	server := server.New(*dir)
	go server.AutoClean()
	log.Println("Starting server...")
	server.HTTP.ListenAndServe()
}
