package main

import (
	"fmt"

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
	var value string

	fmt.Print(prompt)
	fmt.Scanln(&value)

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
