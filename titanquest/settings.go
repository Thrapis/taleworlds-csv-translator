package titanquest

import "tw-translator/translating"

func NewTitanQuestSettings() *translating.TranslationSettings {
	return &translating.TranslationSettings{
		Analyse:                    Analyse,
		PartialStringGetTypeString: PartialStringGetTypeString,
		PartialStringString:        PartialStringString,
		Delimeter:                  "=",
	}
}
