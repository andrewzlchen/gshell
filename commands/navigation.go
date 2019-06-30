package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis"
	"github.com/xchenny/gshell/configs"
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

// CurrentWD returns the current working directory as a string
func CurrentWD() string {
	dir, err := os.Getwd()
	if err != nil {
		return "Could not get current working directory\n"
	}
	return dir + "\n"
}

// ChangeDir changes the current working directory to specified path
func ChangeDir(dir string) string {
	err := os.Chdir(dir)
	if err != nil {
		return fmt.Sprintf("Could not change directory to: %v\n", dir)
	}
	return ""
}

// Wd looks up if the passed in name exists in the database. If it does,
// it will jump to that location in the file system
func Wd(dirName string) string {
	client := redis.NewClient(configs.RedisOptions())
	path, err := client.Get(dirName).Result()
	if err == redis.Nil {
		return fmt.Sprintf("No path associated with the name: '%s'\n", dirName)
	} else if err != nil {
		return fmt.Sprintf("Error occurred while getting filepath: %v\n", err)
	}
	ChangeDir(path)
	return ""
}
