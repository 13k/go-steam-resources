package main

//go:generate ./generate.sh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type genConfig struct {
	BaseImportPath string       `json:"baseImportPath"`
	InputRelPath   string       `json:"inputRelPath"`
	Packages       []*pbPackage `json:"packages"`

	File       string `json:"-"`
	BasePath   string `json:"-"`
	InputPath  string `json:"-"`
	OutputPath string `json:"-"`

	packagesByName map[string]*pbPackage
	template       *template.Template
}

func newGenConfig(basePath string) (*genConfig, error) {
	configFile := filepath.Join(basePath, "config.json")
	config := &genConfig{
		File:       configFile,
		BasePath:   basePath,
		OutputPath: basePath,
	}

	configData, err := ioutil.ReadFile(configFile)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(configData, config); err != nil {
		return nil, err
	}

	config.InputPath = filepath.Join(config.BasePath, config.InputRelPath)

	templateFile := filepath.Join(basePath, "generate.go.template")
	config.template, err = template.ParseFiles(templateFile)

	if err != nil {
		return nil, err
	}

	config.packagesByName = make(map[string]*pbPackage)

	for _, pkg := range config.Packages {
		config.packagesByName[pkg.Name] = pkg
	}

	return config, nil
}

type pbPackage struct {
	Name         string   `json:"name"`
	Dir          string   `json:"directory"`
	Dependencies []string `json:"dependencies"`
	Files        []string `json:"files"`

	config        *genConfig
	relFiles      []string
	deps          []*pbPackage
	importPath    string
	inputPath     string
	inputRelPath  string
	outputPath    string
	outputRelPath string
	args          []string
}

func (p *pbPackage) RelFiles() []string {
	if p.relFiles == nil {
		for _, file := range p.Files {
			p.relFiles = append(p.relFiles, filepath.Join(p.InputRelPath(), file))
		}
	}

	return p.relFiles
}

func (p *pbPackage) PackageDependencies() []*pbPackage {
	if p.deps == nil {
		for _, dep := range p.Dependencies {
			if pkg, ok := p.config.packagesByName[dep]; ok {
				p.deps = append(p.deps, pkg)
			}
		}
	}

	return p.deps
}

func (p *pbPackage) IncludePath() string {
	return p.InputRelPath()
}

func (p *pbPackage) ImportPath() string {
	if p.importPath == "" {
		p.importPath = path.Join(p.config.BaseImportPath, p.Name)
	}

	return p.importPath
}

func (p *pbPackage) InputPath() string {
	if p.inputPath == "" {
		p.inputPath = filepath.Join(p.config.InputPath, p.Dir)
	}

	return p.inputPath
}

func (p *pbPackage) InputRelPath() string {
	if p.inputRelPath == "" {
		rel, err := filepath.Rel(p.OutputPath(), p.InputPath())

		if err != nil {
			// This Should Not Happen™
			panic(err)
		}

		p.inputRelPath = rel
	}

	return p.inputRelPath
}

func (p *pbPackage) OutputPath() string {
	if p.outputPath == "" {
		p.outputPath = filepath.Join(p.config.OutputPath, p.Name)
	}

	return p.outputPath
}

/*
func (p *pbPackage) OutputRelPath() string {
	if p.outputRelPath == "" {
		rel, err := filepath.Rel(p.InputPath(), p.OutputPath())

		if err != nil {
			// This Should Not Happen™
			panic(err)
		}

		p.outputRelPath = rel
	}

	return p.outputPath
}
*/

func (p *pbPackage) GenerationArgs() []string {
	if p.args == nil {
		p.args = []string{
			"protoc",
			"-I", p.IncludePath(),
		}

		var goOutOpts []string
		pkgs := append([]*pbPackage{p}, p.PackageDependencies()...)

		for _, pkg := range pkgs {
			for _, file := range pkg.Files {
				option := fmt.Sprintf("M%s=%s", file, pkg.ImportPath())
				goOutOpts = append(goOutOpts, option)
			}
		}

		goOutOpts = append(goOutOpts, fmt.Sprintf("import_path=%s", p.ImportPath()))
		goOutArg := fmt.Sprintf("--go_out=%s:.", strings.Join(goOutOpts, ","))
		p.args = append(p.args, goOutArg)
		p.args = append(p.args, p.RelFiles()...)
	}

	return p.args
}

func (p *pbPackage) Create() error {
	if err := os.MkdirAll(p.OutputPath(), os.ModePerm); err != nil {
		return err
	}

	pkgShortName := path.Base(p.Name)
	pkgFileName := filepath.Join(p.OutputPath(), "generate.go")
	pkgFile, err := os.Create(pkgFileName)

	if err != nil {
		return err
	}

	defer pkgFile.Close()

	tmplValues := map[string]string{
		"Package":        pkgShortName,
		"GenerationArgs": strings.Join(p.GenerationArgs(), " "),
	}

	if err := p.config.template.Execute(pkgFile, tmplValues); err != nil {
		return err
	}

	return nil
}

func main() {
	var basePath string

	if len(os.Args) > 1 {
		basePath = os.Args[1]
	} else {
		if p, err := os.Getwd(); err != nil {
			panic(err)
		} else {
			basePath = p
		}
	}

	config, err := newGenConfig(basePath)

	if err != nil {
		panic(err)
	}

	for _, pkg := range config.Packages {
		pkg.config = config

		if err := pkg.Create(); err != nil {
			panic(err)
		}
	}
}
