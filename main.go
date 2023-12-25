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
	printSomething("123")

	// in add if we speicfy the types as interface{} saying it can take any value
	// then value will be of type interface{} which will cause problems
	// so we use generics
	value := add(1, 2)
	fmt.Println(value +21)
	value1 := add("1", "2")
	fmt.Println(value1 + "21")
}


func add[T int | float64 | string](a, b T) T {
	// we don't need to check all and ust give supported types
	// this can be used while writing libraries
	return a+b
	// aInt, aIsInt := a.(int)
	// bInt, bIsInt := b.(int)

	// if(aIsInt && bIsInt) {
	//  	return aInt + bInt
	//}

	// aFloat, aIsFloat := a.(float64)
	// bFloat, bIsFloat := b.(float64)

	//if(aIsFloat&&bIsFloat) {
	//  	return aFloat + bFloat
	//}
}
func printSomething(value interface{}) {
	typedVal, ok := value.(int)
	// this checkes if the value if of which type
	if (!ok) {
		fmt.Println("error")
		return
	} else {
		typedVal+=1
		fmt.Println(value, typedVal)
		return
	}
	
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