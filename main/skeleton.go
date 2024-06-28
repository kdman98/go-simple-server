package main

import (
	"context"
	"github.com/scharissis/go-server-skeleton/skeleton"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	apiPort := ":" + skeleton.GetOrDefault("API_PORT", "8000")
	apiPrefix := skeleton.GetOrDefault("API_PREFIX", "")

	runServer(apiPort, apiPrefix)
}

func runServer(port, prefix string) {

	var wg sync.WaitGroup
	defer wg.Wait()

	srv := skeleton.NewServer(prefix)
	s := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      srv,
	}

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
		sig := <-sigs
		signal.Stop(sigs)
		wg.Add(1)
		defer wg.Done()
		log.Printf("received %s, shutting down", sig)
		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("error whilst shutting down HTTP server: %v\n", err)
		}
	}()

	log.Printf("Server running (port %s)...\n", port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("error from ListenAndServe: %v\n", err)
	}
}
