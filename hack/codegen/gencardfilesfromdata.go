package main

import (
	"fmt"
	"os"
)

func generateCardFileFromData(destGeneratedFilePath string, varName string, dataFilePath string, tagDataFiles map[string]string) {
	cards := parseCardsFile(dataFilePath)
	for tag, tagFilePath := range tagDataFiles {
		addTags(cards, tag, tagFilePath)
	}
	generatedOutput :=
		fmt.Sprintf(`package generateddata

import (
	"netrunner-erng/deck-builder/datamodel"
	"netrunner-erng/helpers"
)
	
var %s []*datamodel.Card = []*datamodel.Card{
`, varName)

	for _, card := range cards {
		generatedOutput += "{\n"
		generatedOutput += fmt.Sprintf("Title: %#v,\n", card.Title)
		if card.Cost != nil {
			generatedOutput += fmt.Sprintf("Cost: helpers.ToPtr(%d),\n", *card.Cost)
		}
		generatedOutput += fmt.Sprintf("Type: \"%s\",\n", card.Type)
		if card.Strength != nil {
			generatedOutput += fmt.Sprintf("Strength: helpers.ToPtr(%d),\n", *card.Strength)
		}
		if card.Points != nil {
			generatedOutput += fmt.Sprintf("Points: helpers.ToPtr(%d),\n", *card.Points)
		}
		if card.Trash != nil {
			generatedOutput += fmt.Sprintf("Trash: helpers.ToPtr(%d),\n", *card.Trash)
		}
		if card.Subtype != nil {
			generatedOutput += fmt.Sprintf("Subtype: helpers.ToPtr(%#v),\n", *card.Subtype)
		}
		if card.Inf != nil {
			generatedOutput += fmt.Sprintf("Inf: helpers.ToPtr(%d),\n", *card.Inf)
		}
		generatedOutput += fmt.Sprintf("Set: %#v,\n", card.Set)
		if card.Money != nil {
			generatedOutput += fmt.Sprintf("Money: helpers.ToPtr(%#v),\n", *card.Money)
		}
		if card.Tags != nil {
			generatedOutput += "Tags: []string{\n"
			for _, tag := range card.Tags {
				generatedOutput += fmt.Sprintf("%#v,\n", tag)
			}
			generatedOutput += "},\n"

		}
		generatedOutput += "},\n"
	}

	generatedOutput +=
		`}`

	f, err := os.Create(destGeneratedFilePath)
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(generatedOutput))
	if err != nil {
		panic(err)
	}
}
