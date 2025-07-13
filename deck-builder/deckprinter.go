package main

import (
	"fmt"
	"net/http"
	"sort"
	dbdatamodel "netrunner-erng/deck-builder/deckbuilder/datamodel"
	"strings"
)

func printCorpDeck(w http.ResponseWriter, corpDeck dbdatamodel.CorpDeck) {
	fmt.Fprintf(w, "<h3 style=\"color:#e0e0e0;\">Ampere:</h3>\n")
	fmt.Fprintln(w, "<button onclick=\"copyCorp()\">Copy List</button>")

	fmt.Fprintln(w, "<details>")
	fmt.Fprintln(w, "<summary>decklist</summary>")
	fmt.Fprintln(w, "<div class=\"decklist-content\">")
	corpList := ""
	
	// Extract and sort card titles alphabetically
	cards := corpDeck.GetCards()
	var cardTitles []string
	for title := range cards {
		cardTitles = append(cardTitles, title)
	}
	sort.Strings(cardTitles)
	
	for _, title := range cardTitles {
		fmt.Fprintf(w, "1 %s <br>\n", title)
		corpList += fmt.Sprintf("1 %s\\n", title)
	}
	fmt.Fprintln(w, "</div>")
	fmt.Fprintln(w, "</details>")

	fmt.Fprintf(w, `<script>
	function copyCorp() {
		unsecuredCopyToClipboard("%s");
	}
	</script>`, strings.ReplaceAll(corpList, "\"", "\\\""))
}

func printRunnerDeck(w http.ResponseWriter, runnerDeck dbdatamodel.Deck) {
	fmt.Fprintf(w, "<h3 style=\"color:#e0e0e0;\">Nova:</h3>\n")
	fmt.Fprintln(w, "<button onclick=\"copyRunner()\">Copy List</button>")
	fmt.Fprintln(w, "<details>")
	fmt.Fprintln(w, "<summary>decklist</summary>")
	fmt.Fprintln(w, "<div class=\"decklist-content\">")
	runnerList := ""
	
	// Extract and sort card titles alphabetically
	cards := runnerDeck.GetCards()
	var cardTitles []string
	for title := range cards {
		cardTitles = append(cardTitles, title)
	}
	sort.Strings(cardTitles)
	
	for _, title := range cardTitles {
		fmt.Fprintf(w, "1 %s <br>\n", title)
		runnerList += fmt.Sprintf("1 %s\\n", title)
	}
	fmt.Fprintln(w, "</div>")
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
