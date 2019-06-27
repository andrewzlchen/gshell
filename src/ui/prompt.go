package ui

import "fmt"

// Prompt prints out the current directory and a symbol before the shell command cursor
func Prompt() {
	fmt.Print("gshell: ")
}
