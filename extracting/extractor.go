package extracting

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html/charset"
)

const (
	crfl = "\r\n"
	lf   = "\n"
)

func Extract(in io.Reader, out *[]*DataLine, delimeter string) (*Settings, error) {
	bytes, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}

	enc, _, _ := charset.DetermineEncoding(bytes, "")

	utf8Bytes, err := enc.NewDecoder().Bytes(bytes)
	if err != nil {
		return nil, err
	}

	text := string(utf8Bytes)

	lineDelimeter := lf
	if strings.Contains(text, crfl) {
		lineDelimeter = crfl
	}

	settings := &Settings{
		Encoding:      enc,
		LineDelimeter: lineDelimeter,
	}

	lines := strings.Split(text, lineDelimeter)

	*out = make([]*DataLine, 0)

	for _, line := range lines {
		if before, after, ok := strings.Cut(line, delimeter); ok {
			*out = append(*out, &DataLine{
				Key:   before,
				Value: after,
			})
		}
	}
	return settings, nil
}

func Compose(settings *Settings, out io.Writer, in *[]*DataLine, delimeter string) error {
	text := strings.Builder{}

	for _, dataLine := range *in {
		line := fmt.Sprintf("%s%s%s%s", dataLine.Key, delimeter, dataLine.Value, settings.LineDelimeter)
		text.WriteString(line)
	}

	bytes := []byte(text.String())

	fileBytes, err := settings.Encoding.NewEncoder().Bytes(bytes)
	if err != nil {
		return err
	}

	_, err = out.Write(fileBytes)
	return err
}
