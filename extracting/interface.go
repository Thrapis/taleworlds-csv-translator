package extracting

import "io"

type Extractor func(io.Reader, *[]*DataLine, string) (*Settings, error)

type Composer func(*Settings, io.Writer, *[]*DataLine, string) error
