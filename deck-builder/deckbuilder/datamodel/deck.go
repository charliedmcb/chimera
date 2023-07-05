package datamodel

import (
	"errors"
	"netrunner-erng/deck-builder/datamodel"
)

type Deck interface {
	Add(card *datamodel.Card) error
	GetCards() map[string]*datamodel.Card
	Size() int
}

type deck struct {
	cards map[string]*datamodel.Card
}

func NewDeck() *deck {
	return &deck{
		cards: map[string]*datamodel.Card{},
	}
}

func (cd *deck) Add(card *datamodel.Card) error {
	if _, exists := cd.cards[card.Title]; exists {
		return errors.New("card already exists in deck")
	}
	cd.cards[card.Title] = card
	return nil
}

func (cd *deck) GetCards() map[string]*datamodel.Card {
	return cd.cards
}

func (cd *deck) Size() int {
	return len(cd.cards)
}
