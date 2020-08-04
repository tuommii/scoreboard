package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"miikka.xyz/scoreboard/server"
)

const (
	cleanInterval = time.Minute * 20
	editedAlive   = time.Hour * 1
	createdAlive  = time.Hour * 5
)

func main() {
	static := flag.String("static", "./frontend/public", "Path to static dir")
	port := flag.String("port", "8080", "Port to listen")
	mem := flag.String("mem", "./assets/memory.json", "path to memory.json")
	designs := flag.String("designs", "./assets/designs.json", "path to designs")

	flag.Parse()
	s := server.New(*static, *port, *designs)
	go s.AutoClean(cleanInterval, editedAlive, createdAlive)

	go func() {
		log.Println("Starting server...")
		if err := s.HTTP.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	shutdown(s, *mem)
}

func shutdown(s *server.Server, path string) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	s.LoadMemory(path)
	<-signalCh
	s.SaveMemory(path)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	s.HTTP.Shutdown(ctx)
	log.Println("Exiting...")
	os.Exit(0)
}
