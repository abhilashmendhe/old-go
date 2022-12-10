package main

import (
	"Youtube/NicJackson/Microservices2/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"golang.org/x/net/context"
)

func main() {

	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	// injecting the above log to below handlers
	ph := handlers.NewProducts(l)

	// now we will use gorilla mux, using github.com/gorilla/mux
	sm := mux.NewRouter()
	// set  routing path
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)
	// we will manually create our http server.. below is our custom http server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
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
