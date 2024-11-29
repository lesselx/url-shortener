package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	data "url.shortener.agnes/internal/model"
)

type config struct {
	port int
	env  string
	host string
}

type application struct {
	config config
	db     *gorm.DB
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "environment", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.host, "host", "http://localhost", "Host")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open("link.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	data.SetupDatabase(db)

	app := &application{
		config: cfg,
		db:     db,
	}

	log.Printf("Starting server on %s:%d", cfg.host, cfg.port)
	if err := app.routes().Run(fmt.Sprintf(":%d", app.config.port)); err != nil {
		log.Fatal(err)
	}
}
