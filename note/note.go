package note

import (
	"encoding/json"
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

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.title, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, []byte(json), 0644)
}
