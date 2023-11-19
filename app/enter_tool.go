package app

import (
	"bufio"
	"fmt"
	"os"
)

func (a *App) selectTool() int {

	var (
		tools map[int]string
	)

	tools = map[int]string{}
	for _, tool := range a.toolStorage.Tools {
		fmt.Printf("-- %s \n", tool.Title)
		tools[tool.Id] = tool.Title
	}

	fmt.Print("Enter the tool name: ")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			input := scanner.Text()

			for toolId, toolName := range tools {
				if toolName == input {
					return toolId
				}
			}

			// Prompt for more input
			fmt.Println("Error : tool name is wrong")
			fmt.Print("Enter the tool name: ")
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			break
		}
	}
	return (-1)
}
