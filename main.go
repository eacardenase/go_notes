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

func printSomething(value any) { // could also be used any, as it is an alias of interface
	// switch value.(type) {
	// case string:
	// 	fmt.Printf("The type of %v is of string.\n", value)
	// case int:
	// 	fmt.Printf("The type of %v is of int.\n", value)
	// case float64:
	// 	fmt.Printf("The type of %v is of float.\n", value)
	// default:
	// 	fmt.Println("Easy boy, I don't recognize that type!")
	// }

	// value.(int) works like a guard type

	typedVal, ok := value.(int) // ok would yield true if value is of type int

	fmt.Println(typedVal + 1) // 0 would be the null value
	// fmt.Println(value + 1) // error, invalid type

	if ok {
		fmt.Printf("%v is of type int\n", typedVal)

		fmt.Println(typedVal + 1) // can safely use typedVal value
	}
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
