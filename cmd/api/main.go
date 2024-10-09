package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
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
	dsn  string
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
	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("READINGLIST_DB_DSN"), "POSTGRES DSN")
	flag.Parse()

	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		log.Fatal("failed to open database connection", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping db", err)
	}

	app := application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	log.Printf("db connection pool established")

	addr := fmt.Sprintf(":%d", app.config.port)

	fmt.Println("listening on port", addr)

	server := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
