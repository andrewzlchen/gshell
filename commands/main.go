package commands

import (
	"strings"
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
			fileList := ListFiles()
			output := strings.Join(fileList, "\n")
			handlerOutputChan <- output
		default:
			handlerOutputChan <- "Unknown command!\n"
		}
	}
	close(handlerOutputChan)
}
