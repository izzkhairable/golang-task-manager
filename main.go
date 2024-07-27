package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	taskName   string
	taskDesc   string
	taskStatus string
}

func main() {
	var tasks map[string]Task
	tasks = make(map[string]Task)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Select the action for tasks:")

		fmt.Println(`
Enter 1 to View tasks
Enter 2 to Add tasks
Enter 3 to Update tasks
Enter 4 to Delete tasks
	`)

		fmt.Print("> ")
		actionNameStr, _ := reader.ReadString('\n')
		actionNameStr = strings.TrimSpace(actionNameStr)
		actionName, _ := strconv.Atoi(actionNameStr)
		fmt.Println(actionName)

		switch actionName {
		case 1:
			fmt.Println("---------View Tasks--------")
			viewTasks(tasks)
			fmt.Println("---------------------------")

		case 2:
			fmt.Println("-------Add Tasks-------")
			addTask(tasks)
			fmt.Println("-----------------------")

		case 3:
			fmt.Println("-------Update Tasks-------")
			updateTask(tasks)
			fmt.Println("-------------------------")
		case 4:
			fmt.Println("-------Delete Tasks-------")
			deleteTask(tasks)
			fmt.Println("--------------------------")
		default:
			fmt.Println("Invalid input")
		}

	}

}

func addTask(tasks map[string]Task) {
	//fmt.Println(tasks)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the task name:")
	fmt.Print("> ")
	inputTaskName, _ := reader.ReadString('\n')
	inputTaskName = strings.TrimSpace(inputTaskName)

	fmt.Println("Please enter the task description:")
	fmt.Print("> ")
	inputTaskDesc, _ := reader.ReadString('\n')
	inputTaskDesc = strings.TrimSpace(inputTaskDesc)

	fmt.Println(inputTaskName, inputTaskDesc)
	tasks[inputTaskName] = Task{
		taskName:   inputTaskName,
		taskDesc:   inputTaskDesc,
		taskStatus: "to-do",
	}
}

func updateTask(tasks map[string]Task) {
	//fmt.Println(tasks)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Provide the 'Task Name' of the task you want to update:")
	viewTasks(tasks)
	fmt.Print("> ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	fmt.Println("\nProvide the field you want to update:")
	fmt.Print("> ")
	taskField, _ := reader.ReadString('\n')
	taskField = strings.TrimSpace(taskField)
	taskField = convertToCamelCase(taskField)

	fmt.Println("\nProvide the value you want to update:")
	fmt.Print("> ")
	taskVal, _ := reader.ReadString('\n')
	taskVal = strings.TrimSpace(taskVal)

	if task, exists := tasks[taskName]; exists {
		switch taskField {
		case "taskStatus":
			task.taskStatus = taskVal
		case "taskDescription":
			task.taskDesc = taskVal
		case "taskName":
			task.taskName = taskVal
		default:
			fmt.Printf("Field %s is not valid.\n", taskField)
			return
		}
		tasks[taskName] = task
		fmt.Println("\nUpdated Tasks!\n\n")
	} else {
		fmt.Printf("Task with name %s not found.\n", taskName)
	}

}

func deleteTask(tasks map[string]Task) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nProvide the task you want to delete:")
	viewTasks(tasks)
	fmt.Print("> ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	if _, exists := tasks[taskName]; exists {
		delete(tasks, taskName)
		fmt.Println("\nDeleted Tasks!\n\n")
	} else {
		fmt.Printf("Task with key %s not found.\n", taskName)
	}
}

func viewTasks(tasks map[string]Task) {
	if len(tasks) == 0 {
		fmt.Println("No task found!")
	} else {
		for _, task := range tasks {
			fmt.Println("Task name: ", task.taskName)
			fmt.Println("Task Description: ", task.taskDesc)
			fmt.Println("Task Status: ", task.taskStatus)
			fmt.Println("")
		}
	}
}

func convertToCamelCase(input string) string {
	// Split the input string into words
	words := strings.Fields(input)
	if len(words) == 0 {
		return ""
	}
	result := strings.ToLower(words[0])
	for _, word := range words[1:] {
		result += strings.Title(strings.ToLower(word))
	}
	return result
}
