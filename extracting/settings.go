package extracting

import "golang.org/x/text/encoding"

type Settings struct {
	Encoding      encoding.Encoding
	LineDelimeter string
}
