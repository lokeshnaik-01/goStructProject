package note

import (
	"time"
	"errors"
	"fmt"
)
type Note struct {
	title string
	content string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Title is %v\ncontent is %v\n\n\n", note.title, note.content)
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
