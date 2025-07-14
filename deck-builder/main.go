package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"netrunner-erng/deck-builder/deckbuilder"
)

func main() {

	http.HandleFunc("/generate-decks", func(w http.ResponseWriter, r *http.Request) {
		var name string
		var seed int64

		inputSeed := r.URL.Query().Get("seed")
		if inputSeed != "" {
			name, seed = generateDeckNameAndSeedFromInput(inputSeed)
		} else {
			name, seed = generateDeckNameAndSeed()
		}
		rand.Seed(seed)

		fmt.Fprintln(w, "<!DOCTYPE html>")
		fmt.Fprintln(w, "<html>")

		fmt.Fprintln(w, "<head>")
		fmt.Fprintln(w, "<title>Generate Decks</title>")
		fmt.Fprintln(w, `<link rel="stylesheet" href="/static/style.css">`)
		fmt.Fprintln(w, `<link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;500;700&display=swap" rel="stylesheet">`)
		fmt.Fprintln(w, "</head>")

		fmt.Fprintln(w, "<body>")

		fmt.Fprintln(w, `<nav>
    <a href="/">Home</a>
    <a href="/generate-decks">Generate Decks</a>
    <a href="/banlist/corp">Corp Banlist</a>
    <a href="/banlist/runner">Runner Banlist</a>
    <a href="/econcards/corp">Corp Econ-list</a>
    <a href="/econcards/runner">Runner Econ-list</a>
</nav>`)

		fmt.Fprintln(w, `<main class="container">`)
		fmt.Fprintf(w, "<h1>%s</h1>\n", name)

		corpDeck := deckbuilder.MakeCorpDeck()
		printCorpDeck(w, corpDeck)
		runnerDeck := deckbuilder.MakeRunnerDeck()
		printRunnerDeck(w, runnerDeck)

		fmt.Fprintln(w, `<br><form id="seedForm" method="GET" action="/generate-decks" style="margin-bottom: 10px;">
    <input type="text" name="seed" id="seedInput" placeholder="Enter custom seed..." style="padding: 8px; width: 200px; display: block; margin-bottom: 10px; background-color: #2a2a2a; color: #e0e0e0; border: 1px solid #333; border-radius: 5px;">
    <a href="#" class="button" onclick="document.getElementById('seedForm').submit(); return false;">Generate New Decks</a>
</form>`)

		printCopyScript(w)

		fmt.Fprintln(w, `</main>`)
		fmt.Fprintln(w, "</body>")
		fmt.Fprintln(w, "</html>")
	})

	//Pure text api endpoint for generating a runner deck
	http.HandleFunc("/api/plain-text-runner", func(w http.ResponseWriter, r *http.Request) {
		_, seed := generateDeckNameAndSeed()
		rand.Seed(seed)
		runnerDeck := deckbuilder.MakeRunnerDeck()
		for _, card := range runnerDeck.GetCards() {
			fmt.Fprintf(w, "1 %s \n", card.Title)
		}
	})

	//Pure text api endpoint for generating a corp deck
	http.HandleFunc("/api/plain-text-corp", func(w http.ResponseWriter, r *http.Request) {
		_, seed := generateDeckNameAndSeed()
		rand.Seed(seed)
		corpDeck := deckbuilder.MakeCorpDeck()
		for _, card := range corpDeck.GetCards() {
			fmt.Fprintf(w, "1 %s \n", card.Title)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/homepage.html")
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/favicon.ico")
	})

	http.HandleFunc("/econcards/corp", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/econcards-corp.html")
	})

	http.HandleFunc("/econcards/runner", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/econcards-runner.html")
	})

	http.HandleFunc("/banlist/corp", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/banlist-corp.html")
	})

	http.HandleFunc("/banlist/runner", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/banlist-runner.html")
	})

	http.HandleFunc("/static/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/style.css")
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
