package database

import (
	"database/sql"
	"fmt"
	"log"
)

type DbService struct {
	db *sql.DB
}

type Task struct {
	Id   int
	Task string
	Done bool
}

func InitDB() (*DbService, error) {
	var err error

	db, err := sql.Open("pgx", "postgres://postgres@localhost:15432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DbService{db: db}, nil
}

func (d *DbService) Close() error {
	return d.db.Close()
}

func (d *DbService) AddTask(task string) error {
	log.Println(task)

	query := "INSERT INTO tasks (task, done) VALUES ($1, $2)"

	stmt, err := d.db.Prepare(query)

	if err != nil {
		return fmt.Errorf("Query Preparation Failed: %v", err)
	}
	defer stmt.Close()

	_, executeErr := stmt.Exec(task, false)

	if executeErr != nil {
		return fmt.Errorf("Query Execution Failed: %v", err)
	}
	return nil
}

func (d *DbService) GetTasks() ([]Task, error) {

	query := "SELECT id, task, done FROM tasks"

	rows, err := d.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var tasks []Task

	for rows.Next() {
		var todo Task

		rowErr := rows.Scan(&todo.Id, &todo.Task, &todo.Done)

		if rowErr != nil {
			return nil, err
		}

		tasks = append(tasks, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil

}

func (d *DbService) GetTaskByID(id int) (*Task, error) {

	query := "SELECT id, task, done FROM tasks WHERE id = $1"

	var task Task

	row := d.db.QueryRow(query, id)
	err := row.Scan(&task.Id, &task.Task, &task.Done)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no task was found with task %d", id)
		}
		return nil, err
	}

	return &task, nil

}

func (d *DbService) UpdateTaskById(task Task) error {

	query := "UPDATE tasks SET task = $1, done = $2 WHERE id = $3"

	result, err := d.db.Exec(query, task.Task, task.Done, task.Id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows updated")
	} else {
		fmt.Printf("%d row(s) updated\n", rowsAffected)
	}

	return nil

}

func (d *DbService) DeleTaskWithID(id int) error {

	query := "DELETE FROM tasks WHERE id = $1"

	stmt, err := d.db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no task found with id %d", id)
	}

	fmt.Printf("Deleted %d task(s)\n", rowsAffected)
	return nil

}
