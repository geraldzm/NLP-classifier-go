package main

import (
	"io/ioutil"
	"path/filepath"
)

// read file
func ReadFile(path *string) (string, error) {
	fileBytes, err := ioutil.ReadFile(*path)
	if err != nil {
		return "", err
	}
	return string(fileBytes), nil
}

// read the folder names inside a folder
func ReadFolderNames(path *string) ([]string, error) {
	var folderNames []string
	files, err := ioutil.ReadDir(*path)
	if err != nil {
		return folderNames, err
	}
	for _, file := range files {
		if file.IsDir() {
			folderNames = append(folderNames, file.Name())
		}
	}
	return folderNames, nil
}

// read the txt file names inside a folder
func ReadFileNames(path *string) ([]string, error) {
	var fileNames []string
	files, err := ioutil.ReadDir(*path)
	if err != nil {
		return fileNames, err
	}
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".txt" {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, nil
}
