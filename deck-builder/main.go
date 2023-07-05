package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"netrunner-erng/deck-builder/deckbuilder"
)

func main() {

	http.HandleFunc("/new-deck", func(w http.ResponseWriter, r *http.Request) {
		name, seed := generateDeckNameAndSeed()
		rand.Seed(seed)

		fmt.Fprintln(w, "<!DOCTYPE html>")
		fmt.Fprintln(w, "<html>")

		fmt.Fprintln(w, "<head>")
		fmt.Fprintln(w, "<title>Generate Decks</title>")
		fmt.Fprintln(w, "</head>")

		fmt.Fprintln(w, "<body>")

		fmt.Fprintf(w, "<h3>%s</h3>\n", name)
		fmt.Fprintf(w, "\n")
		corpDeck := deckbuilder.MakeCorpDeck()
		printCorpDeck(w, corpDeck)
		fmt.Fprintf(w, "<br>\n")
		runnerDeck := deckbuilder.MakeRunnerDeck()
		printRunnerDeck(w, runnerDeck)

		printCopyScript(w)

		fmt.Fprintln(w, "</html>")
		fmt.Fprintln(w, "</body>")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./homepage.html")
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./favicon.ico")
	})

	http.HandleFunc("/econcards/corp", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./econcards-corp.html")
	})

	http.HandleFunc("/econcards/runner", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./econcards-runner.html")
	})

	http.HandleFunc("/banlist/corp", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./banlist-corp.html")
	})

	http.HandleFunc("/banlist/runner", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./banlist-runner.html")
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
