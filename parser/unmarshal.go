// Copyright 2023 Nukleros

package parser

type Unmarshaler interface {
	UnmarshalMarkerArg(in string) error
}
