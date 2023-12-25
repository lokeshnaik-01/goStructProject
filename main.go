package main

import (
	"fmt"
	"bufio"
	"os"
	//"errors"
	"strings"
	"example.com/structProject/note"
	"example.com/structProject/todo"
)

// both note and todo have save menthod so instead of duplicating the lines we use interface
// it'll call the save method
type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if(err != nil) {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if(err!=nil) {
		fmt.Println(err)
		return
	}

	
	userNote.Display()
	err = saveData(userNote)
	if(err != nil) {
		fmt.Println("error in saving note")
	}

	todo.Display()
	err = saveData(todo)
	if(err != nil) {
		fmt.Println("error in saving todo")
	}
	
}

func saveData(data saver) error{
	err :=data.Save()
	if(err != nil) {
		fmt.Println(err)
		return err
	}
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	
	content := getUserInput("Note content:")
	
	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	value, err :=reader.ReadString('\n')
	if(err != nil) {
		return ""
	}
	// value will have the escape character so we need to trim it
	value = strings.TrimSuffix(value, "\n")
	// in windows line break is created using "\r\n" so remove both
	value = strings.TrimSuffix(value, "\n")
	return value
}