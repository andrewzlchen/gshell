package commands

import (
	"fmt"
	"io/ioutil"
)

// ListFiles lists all of the files/directory names in the current directory
func ListFiles(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{fmt.Sprintf("Could not read the directory: %v\n", path)}
	}
	var fileList []string
	for _, file := range files {
		fileName := file.Name()

		// make directory names more clear
		if file.IsDir() {
			fileName += "/"
		}
		fileList = append(fileList, fileName)
	}
	return fileList
}

// func ListFiles(path string) {

// }
