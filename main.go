package main

import (
	"errors"
	"fmt"

	"github.com/eacardenase/go_notes/note"
)

func main() {
	note, err := getNote()

	if err != nil {
		fmt.Println("Note could not be created.")

		return
	}

	fmt.Println(note)
}

func getUserInput(prompt string) (string, error) {
	var value string

	fmt.Print(prompt)
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("invalid input")
	}

	return value, nil
}

func getNote() (*note.Note, error) {
	title, err := getUserInput("Note title: ")

	if err != nil {
		fmt.Println(err)

		return &note.Note{}, err
	}

	content, err := getUserInput("Note content: ")

	if err != nil {
		fmt.Println(err)

		return &note.Note{}, err
	}

	note := note.New(title, content)

	return note, nil
}
