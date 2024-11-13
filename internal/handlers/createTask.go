package handlers

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Kyle-Close/tasks/internal/models"
)

func CreateTask(db *sql.DB, task models.Task) (sql.Result, error) {
	query := "INSERT INTO tasks (id, title, description, completed) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, task.Id, task.Title, task.Description, task.Completed)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	fmt.Println("Successfully created task.")
	return result, nil
}
