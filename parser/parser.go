package parser

import (
	"eg-lsp/token"
)

type Parser struct {
	tkns      []token.Token
	pos       int
	lineNo    int
	functions map[string]int
	Errors    map[int]string
}

func NewParser(tkns []token.Token) Parser {
	return Parser{
		tkns:      tkns,
		pos:       0,
		lineNo:    1,
		functions: map[string]int{},
		Errors:    map[int]string{},
	}
}

func (p *Parser) parseToken() {
	currLine := p.lineNo
	switch p.tkns[p.pos].Type {
	case token.INVALID:
		p.Errors[currLine] = "Invalid syntax"
	case token.NEWLINE, token.COMMENT:
		p.lineNo++
	case token.PLUS:
		p.Errors[currLine] = "Invalid position for operator"
		p.pos++
	case token.PRINT:
		p.parsePrint()
	case token.IDENT:
		if p.pos+1 < len(p.tkns) && p.tkns[p.pos+1].Type == token.LPAREN {
			p.parseFnCall()
		} else {
			p.parseIdent()
		}
	case token.IF:
		p.parseIfElse()
	case token.FOR:
		p.parseFor()
	case token.FN:
		p.parseFn()
	}
}
func (p *Parser) Parse() {
	for p.pos < len(p.tkns) {
		p.parseToken()
		p.pos++
	}
}
