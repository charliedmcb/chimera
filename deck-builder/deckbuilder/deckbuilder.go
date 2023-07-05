package deckbuilder

import (
	"math/rand"
	"netrunner-erng/deck-builder/datamodel"
	dbdatamodel "netrunner-erng/deck-builder/deckbuilder/datamodel"
	"netrunner-erng/deck-builder/filereader"
)

func MakeCorpDeck() dbdatamodel.CorpDeck {
	agendas, nonAgendas, nonAgendaMoneyCards := filereader.GetCorpCards()

	etrIceNames := map[string]bool{
		"Bastion":                 true,
		"Battlement":              true,
		"Capacitor":               true,
		"Ice Wall":                true,
		"Maskirovka":              true,
		"Meru Mati":               true,
		"NEXT Silver":             true,
		"Palisade":                true,
		"Paper Wall":              true,
		"Ping":                    true,
		"Quicksand":               true,
		"Sandstone":               true,
		"Seidr Adaptive Barrier":  true,
		"Self-Adapting Code Wall": true,
		"Spiderweb":               true,
		"Tree Line":               true,
		"Vanilla":                 true,
		"Wall of Static":          true,
		"Wraparound":              true,
		"Afshar":                  true,
		"Enigma":                  true,
		"Hortum":                  true,
		"Magnet":                  true,
		"NEXT Bronze":             true,
		"Quandary":                true,
		"Rainbow":                 true,
		"Chimera":                 true,
		"Guard":                   true,
	}
	etrIce := []*datamodel.Card{}

	for _, nonAgenda := range nonAgendas {
		if _, ok := etrIceNames[nonAgenda.Title]; ok {
			etrIce = append(etrIce, nonAgenda)
		}
	}

	corpDeck := dbdatamodel.NewCorpDeck()

	numNonAgendaMoneyCardsPicked := 0
	numEtrIceCardsPicked := 0

	for corpDeck.Size() < 49 {
		if corpDeck.GetAgendaPoints() < 20 {
			nextAgendaInt := rand.Intn(len(agendas))
			nextAgenda := agendas[nextAgendaInt]
			if err := corpDeck.Add(nextAgenda); err != nil {
				continue
			}
			continue
		}

		if numNonAgendaMoneyCardsPicked < 9 {
			nextNonAgendaMoneyCardInt := rand.Intn(len(nonAgendaMoneyCards))
			nextNonAgendaMoneyCard := nonAgendaMoneyCards[nextNonAgendaMoneyCardInt]
			if err := corpDeck.Add(nextNonAgendaMoneyCard); err != nil {
				continue
			}
			numNonAgendaMoneyCardsPicked += 1
			continue
		}

		if numEtrIceCardsPicked < 3 {
			nextEtrIceInt := rand.Intn(len(etrIce))
			nextEtrIce := etrIce[nextEtrIceInt]
			if err := corpDeck.Add(nextEtrIce); err != nil {
				continue
			}
			numEtrIceCardsPicked += 1
			continue
		}

		nextNonAgendaInt := rand.Intn(len(nonAgendas))
		nextNonAgenda := nonAgendas[nextNonAgendaInt]
		if err := corpDeck.Add(nextNonAgenda); err != nil {
			continue
		}
	}

	return corpDeck
}

func MakeRunnerDeck() dbdatamodel.Deck {
	filteredCards, moneyCards := filereader.GetRunnerCards()

	runnerDeck := dbdatamodel.NewDeck()
	numEconCardsPicked := 0

	for _, card := range filteredCards {
		if card.Title == "Peacock" || card.Title == "Aurora" || card.Title == "Creeper" {
			runnerDeck.Add(card)
		}
	}

	for runnerDeck.Size() < 40 {
		if numEconCardsPicked < 8 {
			nextMoneyCardInt := rand.Intn(len(moneyCards))
			nextMoneyCard := moneyCards[nextMoneyCardInt]
			if err := runnerDeck.Add(nextMoneyCard); err != nil {
				continue
			}
			numEconCardsPicked += 1
			continue
		}

		nextCardInt := rand.Intn(len(filteredCards))
		nextCard := filteredCards[nextCardInt]
		if err := runnerDeck.Add(nextCard); err != nil {
			continue
		}
	}

	return runnerDeck
}
