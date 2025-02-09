package translating

type PartialString struct {
	Parts []*StringPart
}

type StringPart struct {
	Type int
	// String, Variable
	Value string
	// Gender, Ternary
	Parts []*StringPart
}
