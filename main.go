package main

import (
	"fmt"
	"math"
	"sync"
)

// n-gram size
var nGram int = 2

// folder path
var folderPath string = "./data/"

func main() {

	// read the folder names inside the folder
	folderNames, err := ReadFolderNames(&folderPath)
	if err != nil {
		panic(err)
	}

	// create a model for each folder
	models := make([]*Model, len(folderNames))
	for i, folderName := range folderNames {
		models[i] = NewModel(folderName)
	}

	// train the models
	wg := sync.WaitGroup{}
	wg.Add(len(models))
	for _, model := range models {
		go model.TrainModelWithFolder(folderPath+model.name+"/", &wg)
	}

	wg.Wait()

	// train a model with the text "this is a mystery text in some lenguaje"
	// and compare it with the other models
	mysteryModel := NewModel("mystery")
	mysteryModel.TrainModelWithText("Voglio andare a studiare")

	for _, model := range models {
		// radians
		r := mysteryModel.CalculateVectorDistance(model)

		// to degrees
		d := r * 180 / math.Pi

		fmt.Printf("Category %s: %f degrees, %f radians\n", model.name, d, r)
	}
}
