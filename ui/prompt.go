package ui

import (
	"fmt"
	"strings"

	"github.com/xchenny/gshell/commands"
)

// Prompt prints out the current directory and a symbol before the shell command cursor
func Prompt() {
	currDir := strings.Trim(commands.CurrentWD(), "\n")
	fmt.Printf("%v gshell: ", currDir)
}
