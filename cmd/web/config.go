package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/CP-Payne/blog_aggregator/internal/database"
	"github.com/CP-Payne/blog_aggregator/pkg/helper"
	"github.com/joho/godotenv"
)

type application struct {
	util *helper.Util
	port string
	DB   *database.Queries
}

func loadConfig() *application {
	godotenv.Load()

	port := os.Getenv("PORT")
	dsn := os.Getenv("DSN")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	utils := helper.NewUtil(errorLog, infoLog)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		errorLog.Fatal("Failed to open connection to database")
	}

	dbQueries := database.New(db)

	app := &application{
		util: utils,
		port: port,
		DB:   dbQueries,
	}

	return app
}
