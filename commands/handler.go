package commands

import (
	"strings"

	"github.com/mitchellh/go-homedir"
)

// CommandHandler is the main function within the 'commands' package that
// will handle user input and map it to a command
func CommandHandler(userInputChan <-chan string, handlerOutputChan chan<- string) {
	for input := range userInputChan {
		toks := strings.Fields(input)
		if len(toks) == 0 {
			continue
		}
		switch toks[0] {
		case "hi":
			handlerOutputChan <- "Hello back!\n"
		case "ls":
			var fileList []string
			if len(toks) > 1 {
				fileList = ListFiles(toks[1])
			} else {
				fileList = ListFiles(".")
			}
			for _, file := range fileList {
				handlerOutputChan <- file
			}
		case "cd":
			var output string
			if len(toks) > 1 {
				output = ChangeDir(toks[1])
			} else {
				home, err := homedir.Dir()
				if err != nil {
					handlerOutputChan <- err.Error()
					continue
				}
				output = ChangeDir(home)
			}
			handlerOutputChan <- output
		case "pwd":
			handlerOutputChan <- CurrentWD()
		default:
			handlerOutputChan <- "Unknown command!\n"
		}
		// this unlocks mutex so new prompt can be printed
		handlerOutputChan <- "/STOP"
	}
	close(handlerOutputChan)
}
