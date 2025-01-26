package parser

import (
	"eg-lsp/token"
)

func (p *Parser) parseIfBlock() {
	errorLine := p.lineNo
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Unexpected end of if block"
		return
	}
	p.pos++ // Past THEN
	for p.pos < len(p.tkns) && p.tkns[p.pos].Type != token.END {
		if p.pos >= len(p.tkns) {
			p.Errors[errorLine] = "Unterminated if block"
			return
		}
		switch p.tkns[p.pos].Type {
		case token.PRINT:
			p.parsePrint()
		case token.IDENT:
			p.parseIdent()
		case token.INPUT:
			p.parseInput()
		case token.NEWLINE:
			p.lineNo++
			p.pos++
			continue
		case token.IF:
			p.parseIfElse()
		case token.FOR:
			p.parseFor()
		case token.RETURN:
			p.parseReturn()
		case token.ELSE:
			return
		default:
			p.pos++
			continue
		}
		p.pos++
	}
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.END {
		p.Errors[errorLine] = "If block missing 'end' keyword"
	}
}

func (p *Parser) parseConditionalBlock() {
	errorLine := p.lineNo
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Unexpected end of conditional block"
		return
	}
	p.pos++ // Skip IF/ELSE IF
	p.parseComparison()
	p.pos++ // Look for THEN
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.THEN {
		p.Errors[errorLine] = "Expected 'then' after condition"
		return
	}
	p.parseIfBlock()
}

func (p *Parser) parseIfElse() {
	errorLine := p.lineNo
	p.parseConditionalBlock()
	for p.pos < len(p.tkns) {
		if p.tkns[p.pos].Type == token.NEWLINE {
			p.lineNo++
			p.pos++
			continue
		}
		if p.tkns[p.pos].Type == token.ELSE {
			p.pos++
			if p.pos >= len(p.tkns) {
				p.Errors[errorLine] = "Unexpected end after 'else'"
				return
			}
			if p.tkns[p.pos].Type == token.IF {
				p.parseConditionalBlock()
			} else {
				p.parseIfBlock()
				break
			}
		} else if p.tkns[p.pos].Type == token.END {
			break
		} else {
			p.Errors[errorLine] = "Expected 'else', 'else if' or 'end'"
			return
		}
	}
}
