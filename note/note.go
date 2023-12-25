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
	Title string `json:"title"`
	Content string
	CreatedAt time.Time `json:"created_at"`
	private string
}

func (note Note) Display() {
	fmt.Printf("Title is %v\ncontent is %v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " " , "_")
	fileName = strings.ToLower(fileName) + ".json"

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
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
		private: "This data is not saved in json as it is private",
	}, nil
}
