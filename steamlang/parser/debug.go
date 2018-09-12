package parser

import (
	"fmt"
	"os"
)

const debugEnvVar = "STEAMLANG_PARSER_DEBUG"

var _verbose bool

func init() {
	_verbose = os.Getenv(debugEnvVar) == "1"
}

func debug(msg string, args ...interface{}) {
	if !_verbose {
		return
	}

	fmt.Printf(msg, args...)
}
