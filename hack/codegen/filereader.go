package main

import (
	"encoding/json"
	"io/ioutil"
	"netrunner-erng/deck-builder/datamodel"
	"os"
)

func parseCardsFile(fileName string) []*datamodel.Card {
	// Open the JSON file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Parse the JSON content
	var cards []*datamodel.Card
	err = json.Unmarshal(content, &cards)
	if err != nil {
		panic(err)
	}
	return cards
}
