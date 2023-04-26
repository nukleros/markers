// Copyright 2023 Nukleros

package inspect

import (
	"github.com/nukleros/markers/marker"
	"github.com/nukleros/markers/parser"
)

type Inspector struct {
	Registry *marker.Registry
}

func NewInspector(registry *marker.Registry) *Inspector {
	return &Inspector{
		Registry: registry,
	}
}

func (s *Inspector) parse(input string) (results []*parser.Result) {
	p := parser.NewParser(input, s.Registry)

	return p.Parse()
}