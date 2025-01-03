package translating

import (
	"regexp"
)

func StringToPartial(text string) *PartialString {
	result := &PartialString{
		Parts: make([]*StringPart, 0),
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

func DetectPart(text string) *StringPart {
	var (
		variableRegex = regexp.MustCompile(variablePattern)
		genderRegex   = regexp.MustCompile(genderPattern)
		ternaryRegex  = regexp.MustCompile(ternaryPattern)
	)

	if match := variableRegex.MatchString(text); match {
		return &StringPart{
			Type:  TypeVariable,
			Value: variableRegex.FindStringSubmatch(text)[1],
		}
	} else if match := genderRegex.MatchString(text); match {
		groups := genderRegex.FindStringSubmatch(text)
		return &StringPart{
			Type: TypeGender,
			Parts: []*StringPart{
				DetectPart(groups[1]),
				DetectPart(groups[2]),
			},
		}
	} else if match := ternaryRegex.MatchString(text); match {
		groups := ternaryRegex.FindStringSubmatch(text)
		return &StringPart{
			Type: TypeTernary,
			Parts: []*StringPart{
				{
					Type:  TypeVariable,
					Value: groups[1],
				},
				DetectPart(groups[2]),
				DetectPart(groups[3]),
			},
		}
	}
	return &StringPart{
		Type:  TypeString,
		Value: text,
	}
}
