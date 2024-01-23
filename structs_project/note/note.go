package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"encoding/json"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Title is %vand its content is below\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err!= nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644 )

}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Please provide required values")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}
