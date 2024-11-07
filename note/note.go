package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("Your note title '%v' has the following content:\n\n%v\n", n.title, n.content)
}

func (n Note) Save() {
	fileName := strings.ReplaceAll(n.title, " ", "")
	fileName = strings.ToLower(fileName)

	os.WriteFile(fileName, []byte(n.content), 0644)
}
