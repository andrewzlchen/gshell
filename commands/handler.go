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
			var fileList []string
			if len(toks) > 1 {
				fileList = ListFiles(toks[1])
			} else {
				fileList = ListFiles(".")
			}
			output := strings.Join(fileList, "\n")
			output += "\n"
			handlerOutputChan <- output
		case "pwd":
			handlerOutputChan <- CurrentWD()
		default:
			handlerOutputChan <- "Unknown command!\n"
		}
	}
	close(handlerOutputChan)
}
