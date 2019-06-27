package commands

import (
	"io/ioutil"
)

// ListFiles lists all of the files/directory names in the current directory
func ListFiles() []string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return []string{"Could not read current directory!\n"}
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
