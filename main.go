package main

import (
	"context"
	"handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"version"
)

// How to try it: PORT=8000 go run main.go
func main() {
	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	port := os.Getenv("PORT")
	port = "8000"
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := handlers.Router(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	log.Print("The service is ready to listen and serve.")

	killSignal := <-interrupt
	switch killSignal {
	case os.Kill:
		log.Print("Got SIGKILL...")
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")

}
