package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/xchenny/gshell/commands"
	"github.com/xchenny/gshell/ui"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	done := false
	userInputChan := make(chan string)
	handlerOutputChan := make(chan string)
	var writeMutex sync.Mutex

	// worker that processes commands
	go commands.CommandHandler(userInputChan, handlerOutputChan)

	// worker that displays results of commands
	go func() {
		for output := range handlerOutputChan {
			fmt.Println(output)
			writeMutex.Unlock()
		}
	}()

	// parent thread parses user input
	for !done {
		writeMutex.Lock()
		ui.Prompt()
		text, _ := reader.ReadString('\n')
		// if we're done, then quit the loop
		if strings.TrimSuffix(text, "\n") == "exit" {
			done = true
			close(userInputChan)
		} else {
			userInputChan <- text
		}
	}
	fmt.Println("Done! See you next time!")
}
