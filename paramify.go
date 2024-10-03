package main

import (
	"bytes"
	"flag"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	_ "embed"

	"github.com/iancoleman/strcase"
	"github.com/yukinagae/paramify/parser"
)

var (
	typeNames    = flag.String("type", "", "comma-separated list of type names; must be set")
	outputPrefix = flag.String("prefix", "", "prefix to be added to the output file")
	outputSuffix = flag.String("suffix", "_gen", "suffix to be added to the output file")
)

func main() {
	flag.Parse()
	if len(*typeNames) == 0 {
		log.Fatalf("the flag -type must be set")
	}
	types := strings.Split(*typeNames, ",")

	dir := "."
	if args := flag.Args(); len(args) == 1 {
		dir = args[0]
	} else if len(args) > 1 {
		log.Fatalf("only one directory at a time")
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("unable to determine absolute filepath for requested path %s: %v",
			dir, err)
	}

	pkg, err := parser.ParsePackage(dir)
	if err != nil {
		log.Fatalf("parsing package in %s: %v", dir, err)
	}

	for _, typeName := range types {
		fields, err := parser.ValuesOfType(pkg, typeName)
		if err != nil {
			log.Fatalf("finding values for type %v: %v", typeName, err)
		}

		var analysis = struct {
			Command        string
			PackageName    string
			TypesAndValues map[string]*parser.Fields
		}{
			Command:     strings.Join(os.Args[1:], " "),
			PackageName: pkg.Name,
			TypesAndValues: map[string]*parser.Fields{
				typeName: fields,
			},
		}

		var buf bytes.Buffer
		if err := generatedTmpl.Execute(&buf, analysis); err != nil {
			log.Fatalf("generating code: %v", err)
		}

		src, err := format.Source(buf.Bytes())
		if err != nil {
			log.Fatalf("formatting output: %s", err)
		}

		output := strcase.ToSnake(*outputPrefix+typeName+*outputSuffix) + ".go"
		outputPath := filepath.Join(dir, output)
		if err := os.WriteFile(outputPath, src, 0644); err != nil {
			log.Fatalf("writing output: %s", err)
		}
	}
}

//go:embed template.go.tpl
var tpl string

var generatedTmpl = template.Must(
	template.New("generated"). //
					Funcs(
			template.FuncMap{
				"ToLower":    strcase.ToLowerCamel,
				"TrimPrefix": strings.TrimPrefix,
			}). //
		Parse(tpl))
