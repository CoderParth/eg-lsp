package parser

import (
	"eg-lsp/token"
)

func (p *Parser) parseReturn() {
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Unexpected end after return")
		return
	}
	p.pos++ // Skip RETURN token
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Expected value or newline after return")
		return
	}
	if p.tkns[p.pos].Type == token.NEWLINE {
		return // void return
	}
	if !p.isComparableTkn() {
		p.printInvalidMsg("Invalid return value")
		return
	}
	p.parseAddSub()
}
