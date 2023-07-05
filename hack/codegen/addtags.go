package main

import (
	"bufio"
	"fmt"
	"netrunner-erng/deck-builder/datamodel"
	"os"
	"strings"
)

func findCard(cards []*datamodel.Card, title string) (*datamodel.Card, error) {
	for _, card := range cards {
		if card.Title == title {
			return card, nil
		}
	}
	return nil, fmt.Errorf("card %s was not found within the given cards", title)
}

func addTags(cards []*datamodel.Card, tag string, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		title := strings.TrimPrefix(strings.TrimSpace(scanner.Text()), "1 ")
		card, err := findCard(cards, title)
		if err != nil {
			panic(err)
		}
		card.Tags = append(card.Tags, tag)
	}
}
