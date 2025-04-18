package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tejas-2232/students_api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// return data
		w.Write([]byte("Welcome to students API")) //converting string to byte slice & passing to write func
	})

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("server started %s", cfg.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
