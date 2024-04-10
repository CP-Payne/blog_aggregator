package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	port     string
}

func loadConfig() *application {
	godotenv.Load()

	port := os.Getenv("PORT")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		port:     port,
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	return app
}
