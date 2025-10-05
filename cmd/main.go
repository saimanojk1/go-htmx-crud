package main

import (
	"html/template"
	"log"
	"task-management/cmd/api"
	"task-management/services/database"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var tmpl *template.Template
var db *database.DbService

func init() {
	tmpl, _ = template.ParseGlob("templates/*.html")
	var err error
	db, err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	server := api.NewAPIServer("localhost:4000", tmpl, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
