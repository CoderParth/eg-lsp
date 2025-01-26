package parser

import (
	"eg-lsp/token"
)

func (p *Parser) parseMath() {
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Unexpected end in math expression")
		return
	}
	p.parseAddSub()
	p.lineNo++
}

func (p *Parser) parseAddSub() {
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Expected value in expression")
		return
	}
	p.parseMulDiv()
	for p.pos < len(p.tkns) && (p.tkns[p.pos].Type == token.PLUS || p.tkns[p.pos].Type == token.MINUS) {
		p.pos++
		if p.pos >= len(p.tkns) {
			p.printInvalidMsg("Expected value after operator")
			return
		}
		p.parseMulDiv()
	}
}

func (p *Parser) parseLParen() {
	p.pos++
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Unclosed parenthesis")
		return
	}
	p.parseAddSub()
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.RPAREN {
		p.printInvalidMsg("Missing closing parenthesis")
		return
	}
	p.pos++
}

func (p *Parser) parseNode() {
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Unexpected end in expression")
		return
	}
	switch p.tkns[p.pos].Type {
	case token.LPAREN:
		p.parseLParen()
	case token.IDENT:
		if p.pos+1 < len(p.tkns) && p.tkns[p.pos+1].Type == token.LPAREN {
			p.parseFnCall()
		} else {
			p.pos++
		}
	case token.NUM, token.STRING, token.TRUE, token.FALSE:
		p.pos++
	default:
		p.printInvalidMsg("Invalid value in expression")
	}
}

func (p *Parser) parseMulDiv() {
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Expected value in expression")
		return
	}
	p.checkIfInvalid()
	p.parseNode()
	for p.pos < len(p.tkns) && (p.tkns[p.pos].Type == token.ASTERISK || p.tkns[p.pos].Type == token.SLASH) {
		p.pos++
		if p.pos >= len(p.tkns) {
			p.printInvalidMsg("Expected value after operator")
			return
		}
		p.parseNode()
		p.checkIfInvalid()
	}
}
