package note

import (
	"time"
	"errors"
	"fmt"
	"os"
	"strings"
	"encoding/json"
)
type Note struct {
	title string
	content string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Title is %v\ncontent is %v\n\n\n", note.title, note.content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " " , "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(note)
	if(err != nil) {
		return err
	}
	os.WriteFile(fileName, json, 0644)
	return nil
}

func New(title, content string) (Note, error){

	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}
