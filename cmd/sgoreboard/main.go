package main

import (
	"flag"
	"log"

	"miikka.xyz/sgoreboard/server"
)

func main() {
	dir := flag.String("dir", "public", "Path to static dir")
	flag.Parse()
	log.Println("Server started...")
	server.Start(*dir)
}
