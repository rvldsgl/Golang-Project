package cmd

import (
	"fmt"
)

func Add(task string) string {
	if task == "" {
		fmt.Println("Error: Argument is invalid")
	}
	addJson(task)
	return fmt.Sprintf("task %s added successfully", task)
}

func Update(id string, task string) string{
	if id == "" || task =="" {
		fmt.Println("Error: Argument is invalid")
	}
	updateJson(id, task)
	return fmt.Sprintf("task with id %s update succesfully", id)
}

func MarkInProgress(id string) string{
	if id == "" {
		fmt.Println("Error: Argument is invalid")
	}
	updateProgress(id)
	return fmt.Sprintf("task with id %s update succesfully", id)
}

func MarkDone(id string) string{
	if id == "" {
		fmt.Println("Error: Argument is invalid")
	}
	updateDone(id)
	return fmt.Sprintf("task with id %s update succesfully", id)
}