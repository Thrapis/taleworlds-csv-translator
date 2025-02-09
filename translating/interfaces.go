package translating

type PartialStringAnalyse func(string) *PartialString

type PartialStringGetTypeString func(*PartialString) []*StringPart

type PartialStringString func(*PartialString) string
