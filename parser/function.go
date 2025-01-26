package parser

import (
	"eg-lsp/token"
)

func (p *Parser) canFindMoreParams() bool {
	if p.pos >= len(p.tkns) {
		return false
	}
	currToken := p.tkns[p.pos].Type
	return currToken == token.IDENT || currToken == token.COMMA
}

func (p *Parser) parseFnName() (string, []string) {
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Expected function name")
		return "", nil
	}
	if p.tkns[p.pos].Type != token.IDENT {
		p.printInvalidMsg("Expected function name")
		return "", nil
	}
	fnName := p.tkns[p.pos].Literal
	p.pos++
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Expected '(' after function name")
		return fnName, nil
	}
	if p.tkns[p.pos].Type != token.LPAREN {
		p.printInvalidMsg("Expected '(' after function name")
		return fnName, nil
	}
	p.pos++
	params := []string{}
	for p.pos < len(p.tkns) && p.canFindMoreParams() {
		if p.tkns[p.pos].Type == token.IDENT {
			params = append(params, p.tkns[p.pos].Literal)
		} else if p.tkns[p.pos].Type != token.COMMA {
			p.printInvalidMsg("Invalid parameter in function definition")
			return fnName, params
		}
		p.pos++
	}
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.RPAREN {
		p.printInvalidMsg("Missing ')' after function parameters")
		return fnName, params
	}
	p.pos++
	return fnName, params
}

func (p *Parser) parseFnBlock() {
	if p.pos >= len(p.tkns) {
		p.printInvalidMsg("Unexpected end of function block")
		return
	}
	p.parseBlock()
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.END {
		p.printInvalidMsg("Function missing 'end' keyword")
	}
}

func (p *Parser) parseFn() {
	startLine := p.lineNo
	if p.pos >= len(p.tkns) {
		return
	}
	p.pos++ // Skip FN token
	fnName, params := p.parseFnName()
	if fnName == "" {
		return
	}
	p.functions[fnName] = len(params)
	if p.pos < len(p.tkns) && p.tkns[p.pos].Type == token.NEWLINE {
		p.lineNo++
	}
	p.parseBlock()
	if p.pos >= len(p.tkns) || p.tkns[p.pos].Type != token.END {
		p.Errors[startLine] = "Function missing 'end' keyword"
	}
}
