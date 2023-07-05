package main

import (
	"fmt"
	"net/http"
	dbdatamodel "netrunner-erng/deck-builder/deckbuilder/datamodel"
	"strings"
)

func printCorpDeck(w http.ResponseWriter, corpDeck dbdatamodel.CorpDeck) {
	fmt.Fprintf(w, "<b>Ampere:</b>\n")
	fmt.Fprintln(w, "<button onclick=\"copyCorp()\">Copy List</button>")

	fmt.Fprintln(w, "<details>")
	fmt.Fprintln(w, "<summary>decklist</summary>")
	corpList := ""
	for _, card := range corpDeck.GetCards() {
		fmt.Fprintf(w, "1 %s <br>\n", card.Title)
		corpList += fmt.Sprintf("1 %s\\n", card.Title)
	}
	fmt.Fprintln(w, "</details>")

	fmt.Fprintf(w, `<script>
	function copyCorp() {
		unsecuredCopyToClipboard("%s");
	}
	</script>`, strings.ReplaceAll(corpList, "\"", "\\\""))
}

func printRunnerDeck(w http.ResponseWriter, runnerDeck dbdatamodel.Deck) {
	fmt.Fprintf(w, "<b>Nova:</b>\n")
	fmt.Fprintln(w, "<button onclick=\"copyRunner()\">Copy List</button>")
	fmt.Fprintln(w, "<details>")
	fmt.Fprintln(w, "<summary>decklist</summary>")
	runnerList := ""
	for _, card := range runnerDeck.GetCards() {
		fmt.Fprintf(w, "1 %s <br>\n", card.Title)
		runnerList += fmt.Sprintf("1 %s\\n", card.Title)
	}
	fmt.Fprintln(w, "</details>")

	fmt.Fprintf(w, `<script>
	function copyRunner() {
		unsecuredCopyToClipboard("%s");
	}
	</script>`, strings.ReplaceAll(runnerList, "\"", "\\\""))
}

func printCopyScript(w http.ResponseWriter) {
	fmt.Fprintln(w, `<script>
	function unsecuredCopyToClipboard(text) {
		const textArea = document.createElement("textarea");
		textArea.value = text;
		document.body.appendChild(textArea);
		textArea.focus();
		textArea.select();
		try {
		  document.execCommand('copy');
		} catch (err) {
		  console.error('Unable to copy to clipboard', err);
		}
		document.body.removeChild(textArea);
	  }
	</script>`)
}
