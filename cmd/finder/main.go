package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Oyal2/SwapiLookup/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Create a signal chan to gracefully shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	a, err := app.New()
	if err != nil {
		log.Fatalln(err)
	}

	// Run the app
	go a.Start(ctx)

	<-sig
	cancel()
}
