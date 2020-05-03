package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/13k/go-steam-resources/steamlang/generator"
	"github.com/13k/go-steam-resources/steamlang/parser"
)

const (
	defaultGoPackage = "steamlang"
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

func abort(err error) {
	warn("%s\n", err)
	os.Exit(1)
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] <input>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(
		&optGoPackage,
		"pkg",
		defaultGoPackage,
		"Name of the Go package to generate",
	)

	flag.StringVar(
		&optProtobufPackage,
		"protopkg",
		"",
		"Import path of the generated protobuf package",
	)

	flag.StringVar(
		&optOutput,
		"o",
		"-",
		"Output file. '-' means stdout",
	)

	flag.Parse()
}

func main() {
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	if optProtobufPackage == "" {
		warn("Option protopkg is required.")
		os.Exit(1)
	}

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
	var buf bytes.Buffer

	if optOutput == "-" {
		output = os.Stdout
	} else {
		output = &buf
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
		f, fErr := os.Create(optOutput)

		if fErr != nil {
			abort(fErr)
		}

		defer f.Close()

		if _, err := io.Copy(f, &buf); err != nil {
			abort(err)
		}

		info("Generated code saved to %q\n", optOutput)
	}
}
