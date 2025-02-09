package main

import (
	"tw-translator/taleworld"
	"tw-translator/translating"
)

func main() {
	settings := taleworld.NewTaleWorldSettings()
	settings.SourceFolder = "C:\\wb_en"
	settings.DestinationFolder = "C:\\wb_be"
	settings.SourceLang = "en"
	settings.TargetLang = "be-BY"
	settings.SourceFolderNameReplace = "en"
	settings.TargetFolderNameReplace = "be"

	translating.StartTranslation(settings)
}
