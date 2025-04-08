package main

import (
	"tw-translator/game/thecoffinofandyandleyley"
	"tw-translator/game/thecoffinofandyandleyley/extracting"
	"tw-translator/translating"
)

func main() {
	settings := thecoffinofandyandleyley.NewTheCoffinOfAndyAndLeyleySettings()
	settings.SourceFolder = "sourcePath"
	settings.DestinationFolder = "targetPath"
	settings.SourceLang = "en"
	settings.TargetLang = "be-BY"
	settings.SourceFolderNameReplace = "en"
	settings.TargetFolderNameReplace = "be"
	settings.Exract = extracting.Extract
	settings.Compose = extracting.Compose

	translating.StartTranslation(settings)

	// filepath := "sourceCsv"
	// splitting.SplitInFiles(filepath, "\r\n\r\n\r\n")
}
