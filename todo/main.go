package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
func main() {
	tasks := []string{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter your tasks here: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)


		if input == "quit" {
			break
		}

		tasks = append(tasks, input)
	}
	fmt.Println("The tasks are:")
	for _, task := range tasks {
		fmt.Println(task)
}


}