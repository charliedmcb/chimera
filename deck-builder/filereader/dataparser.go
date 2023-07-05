package filereader

import (
	"netrunner-erng/deck-builder/datamodel"
	"netrunner-erng/deck-builder/generateddata"
)

func GetCorpCards() ([]*datamodel.Card, []*datamodel.Card, []*datamodel.Card) {
	nonBannedCards := filterCardsWithoutTag(generateddata.ZZ_CorpCards, datamodel.Banned)
	agendas := filterCardsByType(nonBannedCards, datamodel.Agenda)
	nonAgendas := filterCardsByType(nonBannedCards, datamodel.Operation, datamodel.Asset, datamodel.Ice, datamodel.Upgrade)
	nonAgendaMoneyCards := filterCardsByTag(nonAgendas, datamodel.Econ)

	return agendas, nonAgendas, nonAgendaMoneyCards
}

func GetRunnerCards() ([]*datamodel.Card, []*datamodel.Card) {
	nonBannedCards := filterCardsWithoutTag(generateddata.ZZ_RunnerCards, datamodel.Banned)
	filteredCards := filterCardsByType(nonBannedCards, datamodel.Event, datamodel.Program, datamodel.Hardware, datamodel.Resource)
	moneyCards := filterCardsByTag(filteredCards, datamodel.Econ)

	return filteredCards, moneyCards
}
