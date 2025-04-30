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
	"github.com/tejas-2232/students_api/internal/http/handlers/student"
	"github.com/tejas-2232/students_api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	//database setup
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0)"))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started %s", slog.String("address", cfg.HTTPServer.Addr))

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
	// passing the time through context.

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		//logging
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
		// slog.String is used to concatenate the error
	}
	slog.Info("Server shutdown successfully")
}
