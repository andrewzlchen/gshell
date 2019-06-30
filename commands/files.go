package commands

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Cat returns the contents of a file
func Cat(filePath string) []string {
	if _, err := os.Stat(filePath); err != nil && !os.IsExist(err) {
		return []string{fmt.Sprintf("The file: '%v' does not exist: %v\n", filePath, err)}
	}

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return []string{fmt.Sprintf("Error opening file: %v\n", err)}
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return []string{fmt.Sprintf("Error reading file stream: %v\n", err)}
	}

	return []string{fmt.Sprintf("%s\n", buf)}
}
