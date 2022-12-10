package main

import (
	"Youtube/NicJackson/Microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/context"
)

func main() {

	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	// injecting the above log to below handlers
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// now we want to regiester our handler
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// we will manually create our http server.. below is our custom http server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// now gracefull shutdown.... if you do db transaction or downloading large file.. we need to do gracefully...
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigch := make(chan os.Signal)
	signal.Notify(sigch, os.Interrupt)
	signal.Notify(sigch, os.Kill)

	sig := <-sigch
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
