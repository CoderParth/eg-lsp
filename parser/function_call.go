package parser

import (
	"eg-lsp/token"
)

func (p *Parser) parseFnArgs() {
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Unexpected end of function arguments")
		return
	}
	for p.pos < len(p.tkns) && p.tkns[p.pos].Type != token.RPAREN {
		if p.tkns[p.pos].Type == token.COMMA {
			p.pos++
			if p.pos >= len(p.tkns) {
				p.printInvalidMsg("Expected argument after comma")
				return
			}
			continue
		}
		switch p.tkns[p.pos].Type {
		case token.NUM, token.LPAREN, token.IDENT:
			p.parseAddSub()
		case token.STRING:
			p.pos++
		default:
			p.printInvalidMsg("Invalid function argument type")
			return
		}
		// Check for comma or closing parenthesis
		if p.pos >= len(p.tkns) {
			p.printInvalidMsg("Unterminated function arguments")
			return
		}
	}
}

func (p *Parser) parseFnCall() {
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Unexpected end of function call")
		return
	}
	fnName := p.tkns[p.pos].Literal
	// Move past function name
	p.pos++
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.LPAREN {
		p.printInvalidMsg("Expected '(' after function name")
		return
	}
	// Move past opening parenthesis
	p.pos++
	p.parseFnArgs()
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.RPAREN {
		p.printInvalidMsg("Missing closing parenthesis in function call")
		return
	}
	// Move past closing parenthesis
	p.pos++
	if _, exists := p.functions[fnName]; !exists {
		p.printInvalidMsg("Call to undefined function '" + fnName + "'")
	}
}
