package main

import (
	"encoding/csv"
	"io"

	"tw-translator/translating"

	"github.com/gocarina/gocsv"
)

func main() {
	settings := translating.TranslationSettings{
		SourceFolder:            "C:/Users/artbe/OneDrive/Рабочий стол/WarbandRU",
		DestinationFolder:       "C:/Users/artbe/OneDrive/Рабочий стол/WarbandBEfromRU",
		SourceLang:              "ru-RU",
		TargetLang:              "be-BY",
		SourceFolderNameReplace: "ru",
		TargetFolderNameReplace: "sl",
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		reader := csv.NewReader(in)
		reader.Comma = '|'
		reader.LazyQuotes = true
		return reader
	})

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = '|'
		return gocsv.NewSafeCSVWriter(writer)
	})

	translating.Translate(&settings)

	// toTest := "As you wish, {sire/my lady}. {reg6?I:{reg7?You:{s11}}} will be the new {reg3?lady:lord} of {s1}."
	// res := translating.StringToPartial(toTest)
	// for _, v := range res.GetTypeString() {
	// 	fmt.Println(v)
	// }

	// fmt.Println(translating.DetectPart("As you wish"))
	// fmt.Println(translating.DetectPart("{s1}"))
	// fmt.Println(translating.DetectPart("{sire/my lady}"))
	// fmt.Println(translating.DetectPart("{reg3?lady:lord}"))
	// fmt.Println(translating.DetectPart("{reg6?I:{reg7?You:{s11}}}"))
}
