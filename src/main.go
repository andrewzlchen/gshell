package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xchenny/gshell/src/commands"
	"github.com/xchenny/gshell/src/ui"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	done := false
	userInputChan := make(chan string)
	handlerOutputChan := make(chan string)

	// worker that processes commands
	go commands.CommandHandler(userInputChan, handlerOutputChan)

	// worker that displays results of commands
	go func() {
		for output := range handlerOutputChan {
			fmt.Println(output)
		}
	}()

	// parent thread parses user input
	for !done {
		ui.Prompt()
		text, _ := reader.ReadString('\n')
		// if we're done, then quit the loop
		if strings.TrimSuffix(text, "\n") == "exit" {
			done = true
			close(userInputChan)
		} else {
			// TODO: IMPLEMENT WAITGROUPS FOR RESULTS TO PRINT BEFORE NEW PROMPT APPEARS
			userInputChan <- text
		}
	}
	fmt.Println("Done! See you next time!")
}
