package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"unicode"
)

type Model struct {
	name      string
	histogram *HashTable
	norm      float64
}

func NewModel(modelName string) *Model {
	return &Model{modelName, NewHashTable(), 0}
}

func (m *Model) Print() {
	fmt.Println("Model: ", m.name)
	fmt.Println("Histogram: ", m.histogram.table)
	fmt.Println("Norm: ", m.norm)
}

func (m *Model) Clear() {
	m.histogram.table = make(map[string]int, 256)
	m.norm = 0
}

func (m *Model) TrainModelWithFile(path *string, wg *sync.WaitGroup) error {

	defer wg.Done()

	var sb strings.Builder

	// read the file
	content, err := ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", path)
		return err
	}

	for _, r := range content {
		if unicode.IsLetter(r) { // if is letter
			sb.WriteRune(r) // add letter
			if sb.Len() >= nGram {
				m.histogram.Add(strings.ToLower(sb.String()))
				sb.Reset()
			}
		} else {
			sb.Reset() // clean
		}
	}

	return nil
}

func (m *Model) TrainModelWithText(text string) {
	var sb strings.Builder

	for _, r := range text {
		if unicode.IsLetter(r) { // if is letter
			sb.WriteRune(r) // add letter
			if sb.Len() >= nGram {
				m.histogram.Add(strings.ToLower(sb.String()))
				sb.Reset()
			}
		} else {
			sb.Reset() // clean
		}
	}

	m.CalculateNorm()
}

func (m *Model) CalculateNorm() {
	// calculate the norm of the model
	// norm = sqrt(sum of all the values squared)
	var sum int = 0
	for _, value := range m.histogram.table {
		sum += value * value
	}

	m.norm = math.Sqrt(float64(sum))
}

func (m *Model) TrainModelWithFolder(folderPath string, wg *sync.WaitGroup) error {
	defer wg.Done()

	// read all the txt file names inside the folder
	fileNames, err := ReadFileNames(&folderPath)
	if err != nil {
		return err
	}

	wgFiles := sync.WaitGroup{}
	wgFiles.Add(len(fileNames))

	// read all the files inside the folder
	for _, fileName := range fileNames {
		filePath := folderPath + fileName
		go m.TrainModelWithFile(&filePath, &wgFiles)
	}

	wgFiles.Wait()

	m.CalculateNorm()

	return nil
}

func (m *Model) CalculateVectorDistance(otherModel *Model) float64 {

	// calculate the distance between two models
	var dotProduct int = 0

	for key, value := range m.histogram.table {
		otherValue := otherModel.histogram.table[key]
		// ignore the keys that are not in the other model
		if otherValue > 0 {
			dotProduct += value * otherValue
		}
	}

	return float64(dotProduct) / (m.norm * otherModel.norm)
}
