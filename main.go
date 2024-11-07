package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eacardenase/go_notes/note"
)

func main() {
	note, err := getNote()

	if err != nil {
		fmt.Println("Note could not be created.")

		return
	}

	note.Display()
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	value, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)

		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}

func getNote() (note.Note, error) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	note, err := note.New(title, content)

	if err != nil {
		return note, err
	}

	return note, nil
}
