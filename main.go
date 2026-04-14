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
	fmt.Println("4 - Remove tasks")
	fmt.Println("0 - Exit")

	var choice int
	fmt.Scanln(&choice)
	return choice
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		var choice = menu()
		switch choice {
		case 1:
			content, _ := os.ReadFile("ez4ence.txt")
			fmt.Println(string(content))
			fmt.Scanln()
		case 3:
			f, _ := os.OpenFile("ez4ence.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			fmt.Println("Please write something:")

			hi, _ := reader.ReadString('\n')
			hi = strings.TrimSpace(hi)

			f.WriteString(hi + "\n")

			fmt.Println("Task written down.")
			f.Close()
		case 4:
			os.Truncate("ez4ence.txt", 0)
		case 0:
			return
		}
	}
}
