package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func generateHTMLListFile(destGeneratedFilePath string, header string, dataFilePath string) {
	file, err := os.Open(dataFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	generatedOutput := ""

	generatedOutput += fmt.Sprintln("<!DOCTYPE html>")
	generatedOutput += fmt.Sprintln("<html>")
	generatedOutput += fmt.Sprintln("<head>")
	generatedOutput += fmt.Sprintf("<title>%s</title>\n", header)
	generatedOutput += fmt.Sprintln("</head>")

	generatedOutput += fmt.Sprintln("<body>")

	generatedOutput += fmt.Sprintf("<h3>%s:</h3>\n", header)

	generatedOutput += fmt.Sprintln("<p>")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cardTitle := strings.TrimPrefix(strings.TrimSpace(scanner.Text()), "1 ")
		generatedOutput += fmt.Sprintf("%s <br>\n", cardTitle)
	}

	generatedOutput += fmt.Sprintln("</p>")

	generatedOutput += fmt.Sprintln("</body>")

	generatedOutput += fmt.Sprintln("</html>")

	f, err := os.Create(destGeneratedFilePath)
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(generatedOutput))
	if err != nil {
		panic(err)
	}
}
