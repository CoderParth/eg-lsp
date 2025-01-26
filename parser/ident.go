package parser

import (
	"eg-lsp/token"
)

func (p *Parser) parseIdentNode() {
	errorLine := p.lineNo
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Expected value after operator"
		return
	}
	switch p.tkns[p.pos].Type {
	case token.STRING:
		p.pos++
	case token.NUM:
		p.parseAddSub()
	case token.LPAREN:
		p.parseLParen()
	case token.IDENT:
		if p.pos+1 < len(p.tkns) && p.tkns[p.pos+1].Type == token.LPAREN {
			p.parseFnCall()
		} else {
			p.pos++
		}
	case token.TRUE, token.FALSE:
		p.parseBool()
	case token.INPUT:
		p.parseInput()
	default:
		p.Errors[errorLine] = "Invalid value in expression"
	}
}

func (p *Parser) parseIdent() {
	errorLine := p.lineNo
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Unexpected end of input"
		return
	}
	p.pos++ // Past identifier
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.EQ {
		p.Errors[errorLine] = "Expected '=' after identifier"
		return
	}
	p.pos++ // Past '='
	p.parseIdentNode()
	// Handle operators
	for p.pos < len(p.tkns) && p.tkns[p.pos].Type != token.NEWLINE {
		if p.tkns[p.pos].Type != token.PLUS &&
			p.tkns[p.pos].Type != token.MINUS &&
			p.tkns[p.pos].Type != token.ASTERISK &&
			p.tkns[p.pos].Type != token.SLASH {
			p.Errors[errorLine] = "Expected operator"
			return
		}
		p.pos++
		p.parseIdentNode()
	}
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.NEWLINE {
		p.Errors[errorLine] = "Expected newline at end of statement"
		return
	}
	p.lineNo++
}
