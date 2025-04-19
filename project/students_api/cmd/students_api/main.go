package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	fmt.Printf("server started %s", cfg.HTTPServer.Addr)

	// create channel
	// creating this because, if we interuppt the cmd, it will show the interupption in cmd
	// like - exit status 0xc000013a
	done := make(chan os.Signal, 1) // 1 because its buffered channel
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// gracefully shutdown
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-done // listening done -channel
	// unti there is no signal in go channel, the code will not go further. We will  be blocked. Main prog will run continously

	slog.Info("Shutting down the server")

	//saving in var ctx(context)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	//logging
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
		// slog.String is used to concatenate the error
	}

	slog.Info("Server shutdown successfully")

}
