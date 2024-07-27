package main

import (
	"booking-app/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c
	l.Println("recieved terminate , shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
