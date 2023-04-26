// Copyright 2023 Nukleros

package parser

type Registry interface {
	Lookup(name string) bool
	GetDefinition(name string) Definition
}
