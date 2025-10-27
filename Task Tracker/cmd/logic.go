package cmd

import (
	"fmt"
)

func Add(tasks string) string {
	if tasks == "" {
		fmt.Println("Error: Argument is invalid")
	}
	return fmt.Sprintf("task %s added successfully", tasks)
}