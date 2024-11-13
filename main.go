package main

import (
	"log"

	"github.com/Kyle-Close/tasks/internal/database"
	"github.com/Kyle-Close/tasks/internal/handlers"
	"github.com/Kyle-Close/tasks/internal/models"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// TODO: Create endpoint where we get the task and then insert it
	task := models.Task{
		Id:          4,
		Title:       "Wash car",
		Description: "Give the corrola a wash",
		Completed:   true,
	}

	result, err := handlers.CreateTask(db, task)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(result)
}
