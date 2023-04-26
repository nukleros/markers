// Copyright 2023 Nukleros

package parser

import "github.com/nukleros/markers/lexer"

func (p *Parser) consumed(lxt lexer.LexemeType) bool {
	if p.peek().Type == lxt {
		p.next()

		return true
	}

	return false
}
