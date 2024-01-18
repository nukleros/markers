// Copyright 2023 Nukleros

package markers

import "github.com/nukleros/markers/marker"

func NewRegistry() *marker.Registry {
	return marker.NewRegistry()
}

func Define(name string, outputType interface{}) (*marker.Definition, error) {
	return marker.Define(name, outputType)
}
