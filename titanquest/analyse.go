package titanquest

import (
	"regexp"

	"tw-translator/translating"
)

func Analyse(text string) *translating.PartialString {
	result := &translating.PartialString{
		Parts: make([]*translating.StringPart, 0),
	}

	from := 0
	openBrackets := 0
	for i, char := range text {
		if char == '{' {
			if openBrackets == 0 {
				if from != i {
					result.Parts = append(result.Parts, DetectPart(text[from:i]))
				}
				from = i
			}
			openBrackets++
		}
		if char == '}' {
			if openBrackets == 1 {
				if from != i {
					result.Parts = append(result.Parts, DetectPart(text[from:i+1]))
				}
				from = i + 1
			}
			openBrackets--
		}
	}
	if from <= len(text)-1 {
		result.Parts = append(result.Parts, DetectPart(text[from:]))
	}

	return result
}

const (
	variablePattern = `^\{([a-zA-Z0-9_а-яА-ЯёЁ ]+)\}$`
	genderPattern   = `^\{([a-zA-Z0-9_а-яА-ЯёЁ ]+)\/([a-zA-Z0-9_а-яА-ЯёЁ ]+)\}$`
	ternaryPattern  = `^\{(\w+)\?(.*?)\:(.*?)\}$`
)

func DetectPart(text string) *translating.StringPart {
	var (
		variableRegex = regexp.MustCompile(variablePattern)
		genderRegex   = regexp.MustCompile(genderPattern)
		ternaryRegex  = regexp.MustCompile(ternaryPattern)
	)

	if match := variableRegex.MatchString(text); match {
		return &translating.StringPart{
			Type:  TypeVariable,
			Value: variableRegex.FindStringSubmatch(text)[1],
		}
	} else if match := genderRegex.MatchString(text); match {
		groups := genderRegex.FindStringSubmatch(text)
		return &translating.StringPart{
			Type: TypeGender,
			Parts: []*translating.StringPart{
				DetectPart(groups[1]),
				DetectPart(groups[2]),
			},
		}
	} else if match := ternaryRegex.MatchString(text); match {
		groups := ternaryRegex.FindStringSubmatch(text)
		return &translating.StringPart{
			Type: TypeTernary,
			Parts: []*translating.StringPart{
				{
					Type:  TypeVariable,
					Value: groups[1],
				},
				DetectPart(groups[2]),
				DetectPart(groups[3]),
			},
		}
	}
	return &translating.StringPart{
		Type:  TypeString,
		Value: text,
	}
}
