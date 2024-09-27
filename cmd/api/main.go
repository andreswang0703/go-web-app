package main

import (
	"flag"
	"fmt"
	"go-web-app/cmd/internal/data"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")

	app := application{
		config: cfg,
		logger: logger,
	}

	addr := fmt.Sprintf(":%d", app.config.port)

	fmt.Println("listening on port", addr)

	server := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
