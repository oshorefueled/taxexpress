package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/oshorefueled/taxexpress/config"
	"github.com/oshorefueled/taxexpress/handlers"
	"net/http"
	"time"
)

func main () {
	fmt.Println("Initiating TaxExpress Server...")
	r := chi.NewRouter()
	handlers.InitHandlers(r)
	server := createCustomServer(r)
	server.Addr = config.PortToServe()
	fmt.Printf("Starting HTTP server on port: %v\n", config.Port)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("failed to start server ==> err : %s", err)
	}
}

func createCustomServer (router *chi.Mux) *http.Server {
	return &http.Server {
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler: router,
	}
}