package main

import (
	"fmt"
	"netrunner-erng/deck-builder/datamodel"
)

const (
	PATH_TO_GITHUB_REPO_ROOT = "./../../"
	PATH_TO_GENERATED_DATA   = PATH_TO_GITHUB_REPO_ROOT + "deck-builder/generateddata/"
	PATH_TO_DATA_FILES       = PATH_TO_GITHUB_REPO_ROOT + "data/"
	PATH_TO_TAGGED_DATA      = PATH_TO_DATA_FILES + "tags/"

	PATH_TO_HTML_FILES = PATH_TO_GITHUB_REPO_ROOT + "deck-builder/"

	CORP_CARDS_FILE_NAME_SUFFIX   = "corp-cards"
	RUNNER_CARDS_FILE_NAME_SUFFIX = "runner-cards"

	GO_FILE_EXTENSION   = ".go"
	JSON_FILE_EXTENSION = ".json"
	TXT_FILE_EXTENSION  = ".txt"

	CORP_GENERATED_CARDS_GO_VAR_NAME   = "ZZ_CorpCards"
	RUNNER_GENERATED_CARDS_GO_VAR_NAME = "ZZ_RunnerCards"
)

func main() {
	callGenerateCardsFileFromDataForGivenSide(CORP_CARDS_FILE_NAME_SUFFIX, CORP_GENERATED_CARDS_GO_VAR_NAME)
	callGenerateCardsFileFromDataForGivenSide(RUNNER_CARDS_FILE_NAME_SUFFIX, RUNNER_GENERATED_CARDS_GO_VAR_NAME)

	generateHTMLFiles()
}

func generateHTMLFiles() {
	generateHTMLListFile(PATH_TO_HTML_FILES+"econcards-corp.html", "Corp Econ Cards", fmt.Sprint(PATH_TO_TAGGED_DATA, "econ/", CORP_CARDS_FILE_NAME_SUFFIX, TXT_FILE_EXTENSION))
	generateHTMLListFile(PATH_TO_HTML_FILES+"econcards-runner.html", "Runner Econ Cards", fmt.Sprint(PATH_TO_TAGGED_DATA, "econ/", RUNNER_CARDS_FILE_NAME_SUFFIX, TXT_FILE_EXTENSION))
	generateHTMLListFile(PATH_TO_HTML_FILES+"banlist-corp.html", "Corp Banlist Cards", fmt.Sprint(PATH_TO_TAGGED_DATA, "banned/", CORP_CARDS_FILE_NAME_SUFFIX, TXT_FILE_EXTENSION))
	generateHTMLListFile(PATH_TO_HTML_FILES+"banlist-runner.html", "Runner Banlist Cards", fmt.Sprint(PATH_TO_TAGGED_DATA, "banned/", RUNNER_CARDS_FILE_NAME_SUFFIX, TXT_FILE_EXTENSION))
}

func callGenerateCardsFileFromDataForGivenSide(fileNameForSide string, generatedCardsGoVarName string) {
	generatedDataFilePath := fmt.Sprint(PATH_TO_GENERATED_DATA, fileNameForSide, GO_FILE_EXTENSION)
	jsonDataFilePath := fmt.Sprint(PATH_TO_DATA_FILES, fileNameForSide, JSON_FILE_EXTENSION)

	econCardTagsFilePath := fmt.Sprint(PATH_TO_TAGGED_DATA, "econ/", fileNameForSide, TXT_FILE_EXTENSION)
	bannedCardTagsFilePath := fmt.Sprint(PATH_TO_TAGGED_DATA, "banned/", fileNameForSide, TXT_FILE_EXTENSION)

	generateCardFileFromData(
		generatedDataFilePath,
		generatedCardsGoVarName,
		jsonDataFilePath,
		map[string]string{
			datamodel.Econ:   econCardTagsFilePath,
			datamodel.Banned: bannedCardTagsFilePath,
		})
}
