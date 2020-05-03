package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/13k/go-steam-resources/v2/steamlang/generator"
	"github.com/13k/go-steam-resources/v2/steamlang/parser"
)

var (
	flags = &struct {
		GoPackage    string
		ProtoPackage string
		Output       string
		Quiet        bool
	}{
		GoPackage: "steamlang",
		Output:    "-",
	}
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
		&flags.GoPackage,
		"pkg",
		flags.GoPackage,
		"Name of the Go package to generate",
	)

	flag.StringVar(
		&flags.ProtoPackage,
		"protopkg",
		flags.ProtoPackage,
		"Import path of the generated protobuf package",
	)

	flag.StringVar(
		&flags.Output,
		"o",
		flags.Output,
		"Output file. '-' means stdout",
	)

	flag.BoolVar(
		&flags.Quiet,
		"q",
		flags.Quiet,
		"Be quiet",
	)

	flag.Parse()
}

func main() {
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	if flags.ProtoPackage == "" {
		warn("Option protopkg is required.\n")
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

	if flags.Output == "-" {
		output = os.Stdout
	} else {
		output = &buf
	}

	g := &generator.GenGo{
		Package:         flags.GoPackage,
		ProtobufPackage: flags.ProtoPackage,
	}

	err = g.Generate(output, root)

	if err != nil {
		abort(err)
	}

	if flags.Output != "-" {
		f, fErr := os.Create(flags.Output)

		if fErr != nil {
			abort(fErr)
		}

		defer f.Close()

		if _, err := io.Copy(f, &buf); err != nil {
			abort(err)
		}

		if !flags.Quiet {
			info("Generated code saved to %q\n", flags.Output)
		}
	}
}
