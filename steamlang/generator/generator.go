package generator

import (
	"io"

	"github.com/13k/go-steam-resources/steamlang/parser"
)

// Generator is the interface for code generators
type Generator interface {
	Generate(w io.Writer, node *parser.Node) error
}
