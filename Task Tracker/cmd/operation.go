package cmd

import (
	"encoding/json"
	"fmt"
	"time"
	"os"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}



func addJson(desc string) {
	
	// 1. ngecek apakah file sudah ada atau belum
	filename := "data.json"

	jsonData, err := os.ReadFile(filename)

	if err != nil {
		file, err := os.Create(filename)
	
		if err != nil {
			fmt.Println("Error creating JSON file:", err)
			return
		}

		defer file.Close()

		fmt.Println("data.json has been created")
	}

	if len(jsonData) == 0{
		jsonData = []byte("[]")
	}

	// 2. menulis file json

	var tasks []Task

	err = json.Unmarshal(jsonData, &tasks)

	if err != nil{
		fmt.Println("Error unmarshalling JSON", err)
		return
	}

	firstTask := Task{Id: 1, Description: desc, Status: "Todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	tasks = append (tasks, firstTask)

	jsonData, err = json.Marshal(tasks)

	if err != nil{
		fmt.Println("Error writing JSON to file:", err)
		return
	}
  
	os.WriteFile(filename, jsonData, 0644)

	fmt.Println(string(jsonData))


	








}