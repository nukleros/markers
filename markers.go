// Copyright 2023 Nukleros

package markers

import (
	"github.com/nukleros/markers/marker"
	"github.com/nukleros/markers/parser"
)

func NewRegistry() *marker.Registry {
	return marker.NewRegistry()
}

func NewParser(input string, registry parser.Registry) *parser.Parser {
	return parser.NewParser(input, registry)
}

func Define(name string, outputType interface{}) (*marker.Definition, error) {
	return marker.Define(name, outputType)
}
