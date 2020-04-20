package generator

import (
	"bytes"
	"fmt"
	"go/scanner"
	"strings"
)

func NewSyntaxError(code string, errList scanner.ErrorList, contextLines int) error {
	if contextLines == 0 {
		contextLines = 2
	}

	b := &bytes.Buffer{}
	lines := strings.Split(code, "\n")

	for _, e := range errList {
		var i, j int

		if i = e.Pos.Line - contextLines; i < 0 {
			i = 0
		}

		if j = e.Pos.Line + contextLines; j > len(lines) {
			j = len(lines)
		}

		context := lines[i:j]

		for idx, line := range context {
			fmt.Fprintf(b, "%2d | %s\n", idx+i+1, line)
		}

		fmt.Fprintf(b, "\n%s\n\n", e.Error())
	}

	return fmt.Errorf("syntax error:\n%s", b.String())
}
