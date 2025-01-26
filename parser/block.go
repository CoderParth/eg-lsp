package parser

import (
	"eg-lsp/token"
	"fmt"
)

func (p *Parser) parseBlock() {
	errorLine := p.lineNo
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Unexpected end of block"
		return
	}
	for p.pos < len(p.tkns) && p.tkns[p.pos].Type != token.END {
		currToken := p.tkns[p.pos]
		switch currToken.Type {
		case token.PRINT:
			p.parsePrint()
			p.pos++
		case token.IDENT:
			if p.pos+1 < len(p.tkns) && p.tkns[p.pos+1].Type == token.LPAREN {
				p.parseFnCall()
			} else {
				p.parseIdent()
			}
		case token.INPUT:
			p.parseInput()
			p.pos++
		case token.IF:
			p.parseIfElse()
		case token.FOR:
			p.parseFor()
		case token.NEWLINE:
			p.lineNo++
			p.pos++
		case token.RETURN:
			p.parseReturn()
			p.pos++
		default:
			if !p.isValidInBlock(currToken.Type) {
				p.Errors[p.lineNo] = fmt.Sprintf("Unexpected token '%s' in block", currToken.Literal)
			}
			p.pos++
		}
		// Handle newlines after statements
		if p.pos < len(p.tkns) && p.tkns[p.pos].Type == token.NEWLINE {
			p.lineNo++
		}
	}
}

func (p *Parser) isValidInBlock(tokenType string) bool {
	return tokenType == token.COMMENT || tokenType == token.NEWLINE
}
