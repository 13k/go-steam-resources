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

func abort(err error) {
	warn("%s\n", err)
	os.Exit(1)
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
		abort(errf("Usage: %s [options] <input>\n", os.Args[0]))
	}
}

func main() {
	filename := flag.Arg(0)
	f, err := os.Open(filename)

	if err != nil {
		abort(err)
	}

	defer f.Close()

	fset := parser.FileSet{
		filename: f,
	}

	root, err := parser.ParseFile(fset, filename)

	if err != nil {
		abort(err)
	}

	var output io.Writer

	if optOutput == "-" {
		output = os.Stdout
	} else {
		f, fErr := os.Create(optOutput)

		if fErr != nil {
			abort(fErr)
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
		abort(err)
	}

	if optOutput != "-" {
		info("Generated code saved to %q\n", optOutput)
	}
}
