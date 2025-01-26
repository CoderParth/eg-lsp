package parser

import (
	"eg-lsp/token"
)

func (p *Parser) parseBool() {
	errorLine := p.lineNo
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Invalid boolean value"
		return
	}

	switch p.tkns[p.pos].Type {
	case token.TRUE, token.FALSE:
		p.pos++
		return
	default:
		p.Errors[errorLine] = "Expected 'true' or 'false'"
	}
}
