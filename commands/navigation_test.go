package commands

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// MakeTestDir creates a test directory with test files to be used for testing
func MakeDefaultTestDir(testDirPath string) error {
	if _, err := os.Stat(testDirPath); os.IsNotExist(err) {
		err := os.Mkdir(testDirPath, 0755)
		if err != nil {
			return err
		}
	}
	testFiles := []string{"test1.txt", "test2.txt", "test3.txt"}
	for _, file := range testFiles {
		testFilePath := filepath.Join(testDirPath, file)
		_, err := os.Create(testFilePath)
		if err != nil {
			return err
		}
	}
	return nil
}

// Removes Test Directory
func RemoveTestDir(testDirPath string) error {
	err := os.RemoveAll(testDirPath)
	if err != nil {
		return err
	}
	err = os.Remove(testDirPath)
	if err != nil {
		return err
	}
	return nil
}

func TestListFiles(t *testing.T) {
	testDirName := "testDir"
	err := MakeDefaultTestDir(testDirName)
	if err != nil {
		RemoveTestDir(testDirName)
		t.Errorf("Error happened while creating test files: %v\n", err)
	}
	actualFiles := ListFiles(testDirName)
	sort.Strings(actualFiles)
	expectedFiles := []string{
		"test1.txt",
		"test2.txt",
		"test3.txt",
	}
	if len(actualFiles) != len(expectedFiles) {
		RemoveTestDir(testDirName)
		t.Errorf("\nExpected: %v\nActual: %v\n", expectedFiles, actualFiles)
	}
	for i := 0; i < len(actualFiles); i++ {
		if actualFiles[i] != expectedFiles[i] {
			RemoveTestDir(testDirName)
			t.Errorf("\nExpected: %v\nActual: %v\n", expectedFiles, actualFiles)
		}
	}
	RemoveTestDir(testDirName)
}
