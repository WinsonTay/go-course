package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}
type outputtable interface {
	saver
	Display()
}

func main() {
	var err error
	title, content := getNoteData()
	todoText := getUserInput("Todo: ")
	todoList, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(todoList)

	userNote, error := note.New(title, content)
	if error != nil {
		fmt.Println(error)
		return
	}
	outputData(userNote)

}

func saveData(data saver) error {

	err := data.Save()
	if err != nil {
		fmt.Println("Failed to Save File")
		return err
	}
	printSaveLog(data)
	return nil
}
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text

}
func getNoteData() (string, string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")
	return title, content
}

func printSaveLog(value interface{}) {
	_, isNote := value.(note.Note)
	if isNote {
		fmt.Println("Saving Note Success")
	} else {
		fmt.Println("Saving ToDo Success")
	}
}
