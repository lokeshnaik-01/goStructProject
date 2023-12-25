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

// type displayer interface {
//	 Display()
// }

type outputtable interface {
	saver
	Display()
	// embedded interface
}

// type outputtable interface {
//	 Save() error
//	 Display()
// }
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

	err = outputData(todo)
	
	if(err != nil) {
		fmt.Println("error in saving todo")
	}
	err = outputData(userNote)
	if(err != nil) {
		fmt.Println("error in saving note")
	}
	printSomething("Lokesh Naik")
	printSomething(123)
}

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)
	}
	// interface{} will accept any kind of value not specific
	fmt.Println(value)
}

func outputData(data outputtable) error{
	data.Display()
	return saveData(data)
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