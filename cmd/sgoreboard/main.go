package main

import (
	"flag"
	"log"

	"miikka.xyz/sgoreboard/server"
)

func main() {
	dir := flag.String("dir", "./", "Path to static dir")
	flag.Parse()
	log.Println("Server started...")
	server := server.New(*dir)
	go server.CleanGames()
	server.HTTP.ListenAndServe()
}
