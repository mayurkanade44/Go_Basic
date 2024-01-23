package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/node/note"
)

func main() {
	title := getUserInput("Please provide title: ")
	content := getUserInput("Please provide content: ")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving failed")
	}

	fmt.Println("Data successfully saved")

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
