package api

import (
	"html/template"
	"log"
	"net/http"
	"task-management/services/database"
	"task-management/services/tasks"
)

type APIServer struct {
	addr string
	tmpl *template.Template
	db   *database.DbService
}

func NewAPIServer(addr string, tmpl *template.Template, db *database.DbService) *APIServer {
	return &APIServer{addr: addr, tmpl: tmpl, db: db}
}

func (s *APIServer) Run() error {
	tasksHandler := tasks.NewHandler(s.tmpl, s.db)
	taskRouter := tasksHandler.RegisterRoutes()

	router := http.NewServeMux()
	router.Handle("/", taskRouter)

	log.Println("Starting server on port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
