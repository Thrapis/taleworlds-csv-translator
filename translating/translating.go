package translating

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"tw-translator/extracting"
	"tw-translator/utils"

	translategooglefree "github.com/bas24/googletranslatefree"
)

func StartTranslation(settings *TranslationSettings) {
	folder := getFolder(settings.SourceFolder)

	// folder.PrintDeep()

	translateFolder(folder, settings.DestinationFolder, settings)
}

func getFolder(directory string) *Folder {
	path, name := filepath.Split(directory)
	current := &Folder{
		Name:    name,
		Path:    path,
		Folders: make([]Folder, 0),
		Files:   make([]File, 0),
	}

	entries, _ := os.ReadDir(directory)
	for _, entry := range entries {
		if entry.IsDir() {
			subpath := filepath.Join(directory, entry.Name())
			subFolder := getFolder(subpath)
			current.Folders = append(current.Folders, *subFolder)
		} else {
			current.Files = append(current.Files, File{
				FullName: entry.Name(),
				Path:     directory,
			})
		}
	}
	return current
}

func translateFolder(sourceFolder *Folder, destinationFolder string, settings *TranslationSettings) {
	if _, err := os.Stat(destinationFolder); os.IsNotExist(err) {
		os.MkdirAll(destinationFolder, os.ModePerm)
	}

	for _, file := range sourceFolder.Files {
		// file to read
		readFile, err := os.OpenFile(file.FullPath(), os.O_RDONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer readFile.Close()

		lines := []*extracting.DataLine{}

		// extract lines
		extractionSettings, err := extracting.Extract(readFile, &lines, settings.Delimeter)
		if err != nil {
			panic(err)
		}

		translateLines(lines, settings)

		filePathToWrite := filepath.Join(destinationFolder, file.FullName)

		// file to write
		writeFile, err := os.OpenFile(filePathToWrite, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer writeFile.Close()

		// write lines
		extracting.Compose(extractionSettings, writeFile, &lines, settings.Delimeter)
	}

	for _, folder := range sourceFolder.Folders {
		destinationFolderJoin := folder.Name
		if destinationFolderJoin == settings.SourceFolderNameReplace {
			destinationFolderJoin = settings.TargetFolderNameReplace
		}
		translateFolder(&folder, filepath.Join(destinationFolder, destinationFolderJoin), settings)
	}
}

func translateLines(in_out []*extracting.DataLine, settings *TranslationSettings) {
	for i, line := range in_out {
		partialString := settings.Analyse(line.Value)

		for _, part := range settings.PartialStringGetTypeString(partialString) {
			trimmed := strings.TrimSpace(part.Value)

			if len(trimmed) == 0 {
				continue
			}

			leadingSpaces := utils.CountLeadingSpaces(part.Value)
			finalSpaces := utils.CountFinalSpaces(part.Value)

			isUpper := true
			if len(trimmed) > 0 {
				isUpper = utils.IsUpper([]rune(trimmed)[0])
			}

			translated, err := translategooglefree.Translate(trimmed, settings.SourceLang, settings.TargetLang)
			if err != nil {
				panic(err)
			}

			if isUpper {
				runed := []rune(translated)
				translated = strings.ToUpper(string(runed[0])) + string(runed[1:])
			}

			part.Value = fmt.Sprintf("%s%s%s", strings.Repeat(" ", leadingSpaces), translated, strings.Repeat(" ", finalSpaces))
		}

		translatedValue := settings.PartialStringString(partialString)

		fmt.Printf("%d%% (%d of %d) - Translating: \"%.32s\" - \"%.32s\"\n", i*100/len(in_out), i, len(in_out), line.Value, translatedValue)

		in_out[i].Value = translatedValue
	}
}
