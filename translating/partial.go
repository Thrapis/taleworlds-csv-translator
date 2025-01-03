package translating

import (
	"fmt"
	"strings"
)

type PartialString struct {
	Parts []*StringPart
}

func (ps *PartialString) GetTypeString() []*StringPart {
	result := make([]*StringPart, 0, len(ps.Parts))

	for _, part := range ps.Parts {
		result = append(result, part.GetTypeString()...)
	}

	return result
}

func (ps *PartialString) String() string {
	builder := strings.Builder{}
	for _, part := range ps.Parts {
		builder.WriteString(part.String())
	}
	return builder.String()
}

const (
	TypeString = iota
	TypeVariable
	TypeGender
	TypeTernary
)

type StringPart struct {
	Type int
	// String, Variable
	Value string
	// Gender, Ternary
	Parts []*StringPart
}

func (sp *StringPart) GetTypeString() []*StringPart {
	result := make([]*StringPart, 0)

	if sp.Type == TypeString {
		result = append(result, sp)
	} else if sp.Type == TypeGender {
		result = append(result, sp.Parts...)
	} else if sp.Type == TypeTernary {
		result = append(result, sp.Parts[1].GetTypeString()...)
		result = append(result, sp.Parts[2].GetTypeString()...)
	}

	return result
}

func (sp StringPart) String() string {
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
