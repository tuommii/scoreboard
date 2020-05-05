package main

import (
	"context"
	"flag"
	"log"
	"net/http"
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
	dir := flag.String("dir", "./", "Path to static dir")
	flag.Parse()
	server := server.New(*dir)
	go server.AutoClean(cleanInterval, editedAlive, createdAlive)

	go func() {
		log.Println("Starting server...")
		if err := server.HTTP.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	shutdown(server.HTTP)
}

func shutdown(server *http.Server) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Exiting...")
	os.Exit(0)
}
