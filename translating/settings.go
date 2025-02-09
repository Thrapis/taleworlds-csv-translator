package translating

type TranslationSettings struct {
	Analyse                    PartialStringAnalyse
	PartialStringGetTypeString PartialStringGetTypeString
	PartialStringString        PartialStringString
	Delimeter                  string
	SourceFolder               string
	DestinationFolder          string
	SourceLang                 string
	TargetLang                 string
	SourceFolderNameReplace    string
	TargetFolderNameReplace    string
}
