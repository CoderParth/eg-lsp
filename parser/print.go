package parser

import (
	"eg-lsp/token"
)

func (p *Parser) parsePrint() {
	errorLine := p.lineNo
	p.pos++ // Move past 'print'

	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Expected '(' after print"
		return
	}
	if p.tkns[p.pos].Type != token.LPAREN {
		p.Errors[errorLine] = "Expected '(' after print"
		return
	}

	p.pos++ // Move past '('
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Expected value after print("
		return
	}

	switch p.tkns[p.pos].Type {
	case token.STRING:
		p.pos++
	case token.NUM, token.LPAREN:
		p.parseAddSub()
	case token.IDENT:
		p.pos++
	default:
		p.Errors[errorLine] = "Invalid value in print statement"
		return
	}

	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.RPAREN {
		p.Errors[errorLine] = "Expected ')' after print value"
		return
	}
	p.pos++

	// Update line number if we see newline
	if p.pos < len(p.tkns) && p.tkns[p.pos].Type == token.NEWLINE {
		p.lineNo++
	}
}
