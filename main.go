package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var tasks []string
var scanner = bufio.NewScanner(os.Stdin)

func clear() {
	fmt.Print("\033[H\033[2J")
}
func menu() int {
	clear()
	fmt.Print("Please select one of the options\n\n")
	fmt.Println("1 - Add task")
	fmt.Println("2 - Remove task")
	fmt.Println("3 - Modify tasks")
	fmt.Println("4 - View tasks")
	fmt.Println("5 - Save")
	fmt.Println("0 - Save & Exit")

	var choice int = 0
	fmt.Scanln(&choice)
	return choice
}

func viewTask() {
	clear()
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i])
	}
	fmt.Print("\nPress enter to continue")
	fmt.Scanln()
}
func modifyTask() {
	clear()
	for i := 0; i < len(tasks); i++ {
		fmt.Println(i+1, tasks[i])
	}
	fmt.Print("Please enter task to modify (0 to cancel): ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	if choice > len(tasks) || choice < 0 {
		fmt.Println("bullllllshit")
		return
	}
	fmt.Print("Please enter name of new task (0 to cancel): ")
	scanner.Scan()
	input := scanner.Text()
	if input == "0" {
		return
	}

	tasks[choice-1] = input
}
func addTask() {
	clear()
	fmt.Print("Please enter name of task (0 to cancel): ")
	scanner.Scan()
	input := scanner.Text()
	if input == "0" {
		return
	}
	tasks = append(tasks, input)
}
func removeTask() {
	clear()
	for i := 0; i < len(tasks); i++ {
		fmt.Println(i+1, tasks[i])
	}
	var choice int
	fmt.Println("Enter number of task to delete (-1 to delete ALL, 0 to return): ")
	fmt.Scanln(&choice)
	if choice > len(tasks) || choice < -1 {
		fmt.Println("im on some bullllllshit")
		return
	}
	switch choice {
	case -1:
		tasks = tasks[:0]
		return
	case 0:
		return
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

		case 5:
			clear()
			save()
			fmt.Print("Saved!\n\n")

		case 67:
			data, _ := os.ReadFile("ez4ence.txt")
			content := string(data)
			if content != "" {
				tasks = strings.Split(content, "\n")
			}

		case 0:
			save()
			return
		}
	}
}
