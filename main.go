package main

import (
	"fmt"
	"bufio"
	"os"
	//"errors"
	"strings"
	"example.com/structProject/note"
)


func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if(err!=nil) {
		fmt.Println(err)
		return
	}
	userNote.Display()
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