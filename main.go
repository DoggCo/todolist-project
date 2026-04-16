package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menu() int {
	fmt.Print("Please select one of the options\n\n")
	fmt.Println("1 - View tasks")
	fmt.Println("2 - Modify tasks")
	fmt.Println("3 - Add task")
	fmt.Println("4 - Remove ALL tasks")
	fmt.Println("0 - Exit")

	var choice int = 0
	fmt.Scanln(&choice)
	return choice
}

func deleteTask(tasks []string, index int) []string {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid choice")
		return tasks
	}
	return append(tasks[:index], tasks[index+1:]...)
}

func addSingleTask() {
	reader := bufio.NewReader(os.Stdin)
	f, _ := os.OpenFile("ez4ence.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	fmt.Println("Please write something:")

	hi, _ := reader.ReadString('\n')
	hi = strings.TrimSpace(hi)

	f.WriteString(hi + "\n")

	fmt.Println("Task written down.")
	f.Close()
}

func readTasks() ([]string, error) {
	file, _ := os.Open("ez4ence.txt")

	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	return tasks, scanner.Err()
}
func printTasks(tasks []string) {
	for i, task := range tasks {
		fmt.Printf("%d - %s\n", i+1, task)
	}
}
func modifyTask(tasks []string, index int, newValue string) []string {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid choice")
		return tasks
	}
	tasks[index] = newValue
	return tasks
}
func writeTasks(tasks []string) error {
	file, _ := os.Create("ez4ence.txt")

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		_, err := writer.WriteString(task + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func main() {
	for {
		var choice = menu()
		switch choice {
		case 1:
			content, _ := os.ReadFile("ez4ence.txt")
			fmt.Println(string(content))
			fmt.Scanln()
		case 2:
			tasks, _ := readTasks()
			printTasks(tasks)
			fmt.Print("Select task to modify: ")
			var choice int
			fmt.Scanln(&choice)
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter new text: ")
			newValue, _ := reader.ReadString('\n')
			newValue = strings.TrimSpace(newValue)

			tasks = modifyTask(tasks, choice-1, newValue)
			writeTasks(tasks)
		case 3:
			addSingleTask()
		case 4:
			os.Truncate("ez4ence.txt", 0)
		case 0:
			return
		}
	}
}
