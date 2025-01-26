package parser

import (
	"eg-lsp/token"
)

func (p *Parser) parseInput() {
	errorLine := p.lineNo
	p.pos++ // Past 'input'
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Input statement missing '('"
		return
	}
	if p.tkns[p.pos].Type != token.LPAREN {
		p.Errors[errorLine] = "Input statement missing '('"
		return
	}
	p.pos++ // Past '('
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.STRING {
		p.Errors[errorLine] = "Input requires string prompt"
		return
	}
	p.pos++ // Past string
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.RPAREN {
		p.Errors[errorLine] = "Input statement missing ')'"
		return
	}
	p.pos++
	if p.pos < len(p.tkns) && p.tkns[p.pos].Type == token.NEWLINE {
		p.lineNo++
	}
}
