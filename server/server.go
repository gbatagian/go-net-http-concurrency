package server

import (
	"fmt"
	"go-net-http-concurrency/handlers"
	"go-net-http-concurrency/settings"
	"log"
	"net/http"
)

func Run() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server running..."))
	})

	router.HandleFunc("GET /sleep", handlers.Sleep)
	router.HandleFunc("GET /sleep/{n}", handlers.SleepN)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", settings.Host, settings.Port),
		Handler: router,
	}

	log.Printf("Server starts listening at: %s", server.Addr)
	server.ListenAndServe()
}
