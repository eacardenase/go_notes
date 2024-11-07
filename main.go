package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eacardenase/go_notes/note"
	"github.com/eacardenase/go_notes/todo"
)

func main() {
	todo, err := getTodo()

	if err != nil {
		fmt.Println("Todo could not be created.")

		return
	}

	note, err := getNote()

	if err != nil {
		fmt.Println("Note could not be created.")

		return
	}

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("Saving the file failed with error: ", err)

		return
	}

	fmt.Println("Saving the todo succeeded!")

	note.Display()
	err = note.Save()

	if err != nil {
		fmt.Println("Saving the file failed with error: ", err)

		return
	}

	fmt.Println("Saving the note succeeded!")
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

func getTodo() (todo.Todo, error) {
	content := getUserInput("Todo content: ")
	todo, err := todo.New(content)

	if err != nil {
		return todo, err
	}

	return todo, nil
}
