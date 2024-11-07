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

type outputable interface {
	Display()
	saver
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Something")

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

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(note)
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

func outputData(data outputable) error {
	data.Display()

	return saveData(data)
}

func printSomething(value interface{}) { // could also be used any, as it is an alias of interface
	fmt.Print(value)
}
