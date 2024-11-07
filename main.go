package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eacardenase/go_notes/note"
	"github.com/eacardenase/go_notes/todo"
)

type saver interface {
	Save() error
}

func main() {
	note, err := getNote()

	if err != nil {
		fmt.Println("Note could not be created.")

		return
	}

	todo, err := getTodo()

	if err != nil {
		fmt.Println("Todo could not be created.")

		return
	}

	todo.Display()
	err = saveData(todo)

	if err != nil {
		return
	}

	note.Display()
	err = saveData(note)

	if err != nil {
		return
	}
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

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the file failed with error: ", err)

		return err
	}

	fmt.Println("Saving the todo succeeded!")

	return nil
}
