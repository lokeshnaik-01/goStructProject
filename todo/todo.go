package todo

import (
	"errors"
	"fmt"
	"os"
	"encoding/json"
)
type Todo struct {
	Text string
}

func (todo Todo) Display() {
	fmt.Printf("Title is %v\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if(err != nil) {
		return err
	}
	os.WriteFile(fileName, json, 0644)
	return nil
}

func New(text string) (Todo, error){

	if text == ""{
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: text,
	}, nil
}
