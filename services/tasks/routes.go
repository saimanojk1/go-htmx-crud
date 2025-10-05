package tasks

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"task-management/services/database"
)

type Handler struct {
	tmpl *template.Template
	db   *database.DbService
}

func NewHandler(tmpl *template.Template, db *database.DbService) *Handler {
	return &Handler{tmpl: tmpl, db: db}
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", h.homepage)
	router.HandleFunc("GET /gettaskupdateform/{id}", h.getTaskUpdateForm)
	router.HandleFunc("GET /newtaskform", h.getTaskForm)
	router.HandleFunc("GET /tasks", h.fetchTasks)
	router.HandleFunc("POST /tasks", h.addTask)
	router.HandleFunc("PUT /tasks/{id}", h.updateTask)
	router.HandleFunc("POST /tasks/{id}", h.updateTask)
	router.HandleFunc("DELETE /tasks/{id}", h.deleteTask)

	return router
}

func (h *Handler) fetchTasks(w http.ResponseWriter, r *http.Request) {
	todos, _ := h.db.GetTasks()
	fmt.Println(todos)

	err := h.tmpl.ExecuteTemplate(w, "todoList", todos)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) addTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add task called here")
	r.ParseForm()
	task := r.FormValue("task")

	fmt.Println(task)

	err := h.db.AddTask(task)
	if err != nil {
		log.Fatal(err)
	}

	todos, _ := h.db.GetTasks()

	err = h.tmpl.ExecuteTemplate(w, "todoList", todos)
	if err != nil {
		log.Fatal(err)
	}

}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	taskItem := r.FormValue("task")
	//taskStatus, _ := strconv.ParseBool(r.FormValue("done"))
	var taskStatus bool

	fmt.Println(r.FormValue("done"))

	switch strings.ToLower(r.FormValue("done")) {
	case "yes", "on":
		taskStatus = true
	case "no", "off":
		taskStatus = false
	default:
		taskStatus = false
	}

	taskId, _ := strconv.Atoi(r.PathValue("id"))

	task := database.Task{
		Id: taskId, Task: taskItem, Done: taskStatus,
	}

	updateErr := h.db.UpdateTaskById(task)

	if updateErr != nil {
		log.Fatal(updateErr)
	}

	todos, _ := h.db.GetTasks()

	err := h.tmpl.ExecuteTemplate(w, "todoList", todos)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {

	taskId, _ := strconv.Atoi(r.PathValue("id"))

	err := h.db.DeleTaskWithID(taskId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	todos, _ := h.db.GetTasks()

	err = h.tmpl.ExecuteTemplate(w, "todoList", todos)

	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) homepage(w http.ResponseWriter, r *http.Request) {
	err := h.tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func (h *Handler) getTaskForm(w http.ResponseWriter, r *http.Request) {
	err := h.tmpl.ExecuteTemplate(w, "addTaskForm", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) getTaskUpdateForm(w http.ResponseWriter, r *http.Request) {

	taskId, _ := strconv.Atoi(r.PathValue("id"))

	task, err := h.db.GetTaskByID(taskId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = h.tmpl.ExecuteTemplate(w, "updateTaskForm", task)
	if err != nil {
		log.Fatal(err)
	}

}
