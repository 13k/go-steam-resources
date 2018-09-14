package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/13k/go-steam-resources/steamlang/generator"
	"github.com/13k/go-steam-resources/steamlang/parser"
)

const (
	defaultGoPackage       = "steamlang"
	defaultProtobufPackage = "github.com/13k/go-steam-resources/protobuf"
)

var (
	optGoPackage       string
	optProtobufPackage string
	optOutput          string
)

func info(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func warn(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
}

func errf(msg string, args ...interface{}) error {
	return fmt.Errorf(msg, args...)
}

func abort(err error, code int) {
	warn("%s\n", err)
	os.Exit(code)
}

func init() {
	flag.StringVar(&optGoPackage, "pkg", defaultGoPackage,
		"Name of the Go package to generate")
	flag.StringVar(&optProtobufPackage, "protopkg", defaultProtobufPackage,
		"Import path of the generated protobuf package")
	flag.StringVar(&optOutput, "o", "-",
		"Output file. '-' means stdout")

	flag.Parse()

	if len(flag.Args()) < 1 {
		abort(errf("Usage: %s [options] <input>\n", os.Args[0]), 1)
	}
}

func main() {
	root, err := parser.ParseFile(flag.Arg(0))

	if err != nil {
		abort(err, 1)
	}

	var output io.Writer

	if optOutput == "-" {
		output = os.Stdout
	} else {
		f, fErr := os.Create(optOutput)

		if fErr != nil {
			abort(fErr, 1)
		}

		defer f.Close()
		output = f
	}

	g := &generator.GenGo{
		Package:         optGoPackage,
		ProtobufPackage: optProtobufPackage,
	}

	err = g.Generate(output, root)

	if err != nil {
		abort(err, 1)
	}

	if optOutput != "-" {
		info("Generated code saved to %q\n", optOutput)
	}
}
