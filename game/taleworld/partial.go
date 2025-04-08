package taleworld

import (
	"fmt"
	"strings"

	"tw-translator/translating"
)

func PartialStringGetTypeString(ps *translating.PartialString) []*translating.StringPart {
	result := make([]*translating.StringPart, 0, len(ps.Parts))

	for _, part := range ps.Parts {
		result = append(result, StringPartGetTypeString(part)...)
	}

	return result
}

func PartialStringString(ps *translating.PartialString) string {
	builder := strings.Builder{}
	for _, part := range ps.Parts {
		builder.WriteString(StringPartString(part))
	}
	return builder.String()
}

const (
	TypeString = iota
	TypeVariable
	TypeGender
	TypeTernary
)

func StringPartGetTypeString(sp *translating.StringPart) []*translating.StringPart {
	result := make([]*translating.StringPart, 0)

	if sp.Type == TypeString {
		result = append(result, sp)
	} else if sp.Type == TypeGender {
		result = append(result, sp.Parts...)
	} else if sp.Type == TypeTernary {
		result = append(result, StringPartGetTypeString(sp.Parts[1])...)
		result = append(result, StringPartGetTypeString(sp.Parts[2])...)
	}

	return result
}

func StringPartString(sp *translating.StringPart) string {
	switch sp.Type {
	default:
		fallthrough
	case TypeString:
		return sp.Value
	case TypeVariable:
		return fmt.Sprintf("{%s}", sp.Value)
	case TypeGender:
		return fmt.Sprintf("{%v/%v}", sp.Parts[0], sp.Parts[1])
	case TypeTernary:
		return fmt.Sprintf("{%s?%v:%v}", sp.Parts[0].Value, sp.Parts[1], sp.Parts[2])
	}
}
