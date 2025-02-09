package taleworld

import "tw-translator/translating"

func NewTaleWorldSettings() *translating.TranslationSettings {
	return &translating.TranslationSettings{
		Analyse:                    Analyse,
		PartialStringGetTypeString: PartialStringGetTypeString,
		PartialStringString:        PartialStringString,
		Delimeter:                  "|",
	}
}
