package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"bufio"
)

var tasks []string
var scanner = bufio.NewScanner(os.Stdin)

func clear() {
	fmt.Print("\033[H\033[2J")
}
func menu() int {
	fmt.Print("Please select one of the options\n\n")
	fmt.Println("1 - Add task")
	fmt.Println("2 - Remove task")
	fmt.Println("3 - Modify tasks")
	fmt.Println("4 - View task")
	fmt.Println("0 - Save & Exit")

	var choice int = 0
	fmt.Scanln(&choice)
	return choice
}

func viewTask() {
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i])
	}
	fmt.Scanln()
	clear()
}
func modifyTask() {
	for i := 0; i < len(tasks); i++ {
		fmt.Println(i+1, tasks[i])
	}
	var choice int
	fmt.Scanln(&choice)
	var newTask string
	fmt.Scanln(&newTask)

	tasks[choice-1] = newTask
}
func addTask() {
	scanner.Scan()
	input := scanner.Text()
	tasks = append(tasks, input)
}
func removeTask() {
	for i := 0; i < len(tasks); i++ {
		fmt.Println(i+1, tasks[i])
	}
	var choice int
	fmt.Scanln(&choice)
	if choice == -1 {
		tasks = tasks[:0]
	}
	tasks = slices.Delete(tasks, choice-1, choice)
}
func save() {
	content := strings.Join(tasks, "\n")
	os.WriteFile("ez4ence.txt", []byte(content), 0644)
}

func main() {
	data, _ := os.ReadFile("ez4ence.txt")
	content := string(data)
	if content != "" {
		tasks = strings.Split(content, "\n")
	}
	for {
		var choice int = menu()
		switch choice {
		case 1:
			addTask()

		case 2:
			removeTask()

		case 3:
			modifyTask()

		case 4:
			viewTask()

		case 0:
			save()
			return
		}
	}
}
