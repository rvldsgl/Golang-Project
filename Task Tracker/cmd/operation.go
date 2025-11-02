package cmd

import (
	"encoding/json"
	"fmt"
	"time"
	"os"
	"strconv"
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

	firstTask := Task{Id: len(tasks) + 1, Description: desc, Status: "Todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	tasks = append (tasks, firstTask)

	jsonData, err = json.Marshal(tasks)

	if err != nil{
		fmt.Println("Error writing JSON to file:", err)
		return
	}
  
	os.WriteFile(filename, jsonData, 0644)

	fmt.Println(string(jsonData))

}

func updateJson(id string, desc string){
	idInt, err := strconv.Atoi(id)

	filename := "data.json"
	jsonData, err := os.ReadFile(filename)
	var tasks []Task

	err = json.Unmarshal(jsonData, &tasks)

	if err != nil{
		fmt.Println("Error unmarshalling JSON", err)
		return
	}

	for i,t := range tasks{
		if t.Id == idInt{
			tasks[i].Description = desc
			tasks[i].UpdatedAt = time.Now()
		}
	}

	jsonData, err = json.Marshal(tasks)

	if err != nil{
		fmt.Println("Error writing JSON to file:", err)
		return
	}
  
	os.WriteFile(filename, jsonData, 0644)

	fmt.Println(string(jsonData))

}

func updateProgress(id string){
	idInt, err := strconv.Atoi(id)

	filename := "data.json"
	jsonData, err := os.ReadFile(filename)
	var tasks []Task

	err = json.Unmarshal(jsonData, &tasks)

	if err != nil{
		fmt.Println("Error unmarshalling JSON", err)
		return
	}

	for i,t := range tasks{
		if t.Id == idInt{
			tasks[i].Status = "in progress"
			tasks[i].UpdatedAt = time.Now()
		}
	}

	jsonData, err = json.Marshal(tasks)

	if err != nil{
		fmt.Println("Error writing JSON to file:", err)
		return
	}
  
	os.WriteFile(filename, jsonData, 0644)

	fmt.Println(string(jsonData))

}

func updateDone(id string){
	idInt, err := strconv.Atoi(id)

	filename := "data.json"
	jsonData, err := os.ReadFile(filename)
	var tasks []Task

	err = json.Unmarshal(jsonData, &tasks)

	if err != nil{
		fmt.Println("Error unmarshalling JSON", err)
		return
	}

	for i,t := range tasks{
		if t.Id == idInt{
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now()
		}
	}

	jsonData, err = json.Marshal(tasks)

	if err != nil{
		fmt.Println("Error writing JSON to file:", err)
		return
	}
  
	os.WriteFile(filename, jsonData, 0644)

	fmt.Println(string(jsonData))

}

func List(){
	filename := "data.json"
	jsonData, err := os.ReadFile(filename)
	var tasks []Task

	err = json.Unmarshal(jsonData, &tasks)

	if err != nil{
		fmt.Println("Error unmarshalling JSON", err)
		return
	}

	fmt.Println(tasks)

}

func ListDone(){
	filename := "data.json"
	jsonData, err := os.ReadFile(filename)
	var tasks []Task

	err = json.Unmarshal(jsonData, &tasks)

	if err != nil{
		fmt.Println("Error unmarshalling JSON", err)
		return
	}

	var done  []Task
	for i,t := range tasks{
		if t.Status == "done"{
			done = append(done, tasks[i])
		}
	}

	fmt.Println(done)

}

func delete(id string){
	idInt, err := strconv.Atoi(id)

	filename := "data.json"
	jsonData, err := os.ReadFile(filename)
	var tasks []Task

	err = json.Unmarshal(jsonData, &tasks)

	if err != nil{
		fmt.Println("Error unmarshalling JSON", err)
		return
	}

	for i, t := range tasks{
		if t.Id == idInt{
			tasks = append(tasks[:i],tasks[i+1:]...)
		}
	}

	jsonData, err = json.Marshal(tasks)

	if err != nil{
		fmt.Println("Error writing JSON to file:", err)
		return
	}
  
	os.WriteFile(filename, jsonData, 0644)

	fmt.Println(string(jsonData))
	
}