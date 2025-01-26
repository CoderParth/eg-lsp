package parser

import (
	"eg-lsp/token"
)

func (p *Parser) isCompareOperator() bool {
	if p.pos >= len(p.tkns) {
		return false
	}
	currToken := p.tkns[p.pos].Type
	return currToken == token.EQEQ || currToken == token.NOTEQ ||
		currToken == token.LT || currToken == token.LTEQ ||
		currToken == token.GT || currToken == token.GTEQ
}

func (p *Parser) isComparableTkn() bool {
	if p.pos >= len(p.tkns) {
		return false
	}
	currToken := p.tkns[p.pos].Type
	return currToken == token.NUM || currToken == token.LPAREN ||
		currToken == token.IDENT || currToken == token.STRING ||
		currToken == token.TRUE || currToken == token.FALSE
}

func (p *Parser) correctCompTkn() {
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Unexpected end in comparison")
		return
	}
	currTkn := p.tkns[p.pos]
	switch currTkn.Type {
	case token.NUM, token.LPAREN:
		p.parseAddSub()
		if p.pos > 0 {
			p.pos--
		}
	case token.IDENT, token.STRING:
		p.pos++
	case token.TRUE, token.FALSE:
		p.parseBool()
	default:
		p.printInvalidMsg("Invalid comparison value")
	}
}

func (p *Parser) parseSingleComparison() {
	if !p.isComparableTkn() {
		p.printInvalidMsg("Expected value to compare")
		return
	}
	p.correctCompTkn()
	p.pos++

	if p.pos >= len(p.tkns) {
		return
	}
	if p.tkns[p.pos].Type == token.OR || p.tkns[p.pos].Type == token.AND {
		return
	}
	if !p.isCompareOperator() {
		return
	}
	p.pos++
	if !p.isComparableTkn() {
		p.printInvalidMsg("Expected value after comparison operator")
		return
	}
	p.correctCompTkn()
}

func (p *Parser) parseComparison() {
	errorLine := p.lineNo
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Expected comparison"
		return
	}
	// Check first value
	if !p.isComparableTkn() {
		p.Errors[errorLine] = "Expected value before comparison operator"
		return
	}
	p.parseComparisonNode()
}

func (p *Parser) parseComparisonNode() {
	p.parseSingleComparison()
	p.pos++
	if p.pos >= len(p.tkns) {
		return
	}
	// Handle DO token
	if p.tkns[p.pos].Type == token.DO {
		p.pos-- // Back up one token since parseFor expects to find DO
		return
	}
}
