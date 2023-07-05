package filereader

import (
	"netrunner-erng/deck-builder/datamodel"
	"strings"
)

const (
	RWR_SET = "Rebellion Without Rehearsal"
)

func filterCardsByType(cards []*datamodel.Card, typeStrs ...string) []*datamodel.Card {
	filteredCards := []*datamodel.Card{}
	for _, card := range cards {
		for _, typeStr := range typeStrs {
			if card.Type == typeStr {
				filteredCards = append(filteredCards, card)
			}
		}
	}
	return filteredCards
}

func filterCardsBySet(cards []*datamodel.Card, setStrs ...string) []*datamodel.Card {
	filteredCards := []*datamodel.Card{}
	for _, card := range cards {
		for _, setStr := range setStrs {
			if strings.HasPrefix(card.Set, setStr) {
				filteredCards = append(filteredCards, card)
			}
		}
	}
	return filteredCards
}

func filterCardsByTag(cards []*datamodel.Card, tags ...string) []*datamodel.Card {
	filteredCards := []*datamodel.Card{}
	for _, card := range cards {
		for _, tag := range tags {
			for _, cardTag := range card.Tags {
				if cardTag == tag {
					filteredCards = append(filteredCards, card)
				}
			}
		}
	}
	return filteredCards
}

func filterCardsWithoutTag(cards []*datamodel.Card, tags ...string) []*datamodel.Card {
	filteredCards := []*datamodel.Card{}
	for _, card := range cards {
		cardDoesNotHaveAnyTag := true
		for _, tag := range tags {
			for _, cardTag := range card.Tags {
				if cardTag == tag {
					cardDoesNotHaveAnyTag = false
				}
			}
		}
		if cardDoesNotHaveAnyTag {
			filteredCards = append(filteredCards, card)
		}
	}
	return filteredCards
}
