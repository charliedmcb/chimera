package datamodel

import (
	"errors"
	"netrunner-erng/deck-builder/datamodel"
)

type CorpDeck interface {
	Add(card *datamodel.Card) error
	GetAgendaPoints() int
	GetCards() map[string]*datamodel.Card
	Size() int
}

type corpDeck struct {
	*deck
	agendaPoints int
}

func NewCorpDeck() *corpDeck {
	return &corpDeck{
		deck:         NewDeck(),
		agendaPoints: 0,
	}
}

func (cd *corpDeck) Add(card *datamodel.Card) error {
	if card.Type == datamodel.Agenda {
		if *card.Points+cd.agendaPoints > 20 {
			return errors.New("card would put the deck past 20 agenda points")
		}
	}
	if err := cd.deck.Add(card); err != nil {
		return err
	}
	if card.Type == datamodel.Agenda {
		cd.agendaPoints += *card.Points
	}
	return nil
}

func (cd *corpDeck) GetAgendaPoints() int {
	return cd.agendaPoints
}
