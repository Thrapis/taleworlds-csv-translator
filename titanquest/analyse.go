package titanquest

import (
	"regexp"

	"tw-translator/translating"
)

const (
	specialCurvePattern  = `\{([\^\%\.\:\+a-zA-Z0-9_ ]+)\}`
	specialSquarePattern = `\[([\^\%\.\:\+a-zA-Z0-9_ ]+)\]`
	variablePattern      = `(\%[a-zA-Z0-9_])`
)

func Analyse(text string) *translating.PartialString {
	var (
		specialCurveRegex  = regexp.MustCompile(specialCurvePattern)
		specialSquareRegex = regexp.MustCompile(specialSquarePattern)
		variableRegex      = regexp.MustCompile(variablePattern)
	)

	result := &translating.PartialString{
		Parts: make([]*translating.StringPart, 0),
	}

	focusString := text

	for {
		if len(focusString) == 0 {
			break
		}

		fistIndex := len(focusString)
		lastIndex := len(focusString)

		specialCurveIndeces := specialCurveRegex.FindStringIndex(focusString)
		if specialCurveIndeces != nil && specialCurveIndeces[0] < fistIndex {
			fistIndex, lastIndex = specialCurveIndeces[0], specialCurveIndeces[1]
		}

		specialSquareIndeces := specialSquareRegex.FindStringIndex(focusString)
		if specialSquareIndeces != nil && specialSquareIndeces[0] < fistIndex {
			fistIndex, lastIndex = specialSquareIndeces[0], specialSquareIndeces[1]
		}

		variableIndeces := variableRegex.FindStringIndex(focusString)
		if variableIndeces != nil && variableIndeces[0] < fistIndex {
			fistIndex, lastIndex = variableIndeces[0], variableIndeces[1]
		}

		if fistIndex == len(focusString) {
			result.Parts = append(result.Parts, DetectPart(focusString))
			break
		}

		before := focusString[:fistIndex]
		target := focusString[fistIndex:lastIndex]
		after := focusString[lastIndex:]

		if len(before) > 0 {
			result.Parts = append(result.Parts, DetectPart(before))
		}

		result.Parts = append(result.Parts, DetectPart(target))

		if len(after) == 0 {
			break
		}
		focusString = after
	}

	return result
}

func DetectPart(text string) *translating.StringPart {
	var (
		specialCurveRegex  = regexp.MustCompile(specialCurvePattern)
		specialSquareRegex = regexp.MustCompile(specialSquarePattern)
		variableRegex      = regexp.MustCompile(variablePattern)
	)

	if match := specialCurveRegex.MatchString(text); match {
		return &translating.StringPart{
			Type:  TypeSpecialCurve,
			Value: specialCurveRegex.FindStringSubmatch(text)[1],
		}
	} else if match := specialSquareRegex.MatchString(text); match {
		return &translating.StringPart{
			Type:  TypeSpecialSquare,
			Value: specialSquareRegex.FindStringSubmatch(text)[1],
		}
	} else if match := variableRegex.MatchString(text); match {
		return &translating.StringPart{
			Type:  TypeVariable,
			Value: variableRegex.FindStringSubmatch(text)[1],
		}
	}

	return &translating.StringPart{
		Type:  TypeString,
		Value: text,
	}
}
