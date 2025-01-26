package parser

import (
	"eg-lsp/token"
	"fmt"
)

func (p *Parser) printInvalid() {
	errMsg := fmt.Sprintf(`Invalid syntax: %v`, p.tkns[p.pos].Literal)
	p.Errors[p.lineNo] = errMsg
}

func (p *Parser) printInvalidMsg(s string) {
	errMsg := fmt.Sprintf(`%v`, s)
	p.Errors[p.lineNo] = errMsg
}

func (p *Parser) checkIfInvalid() {
	if p.tkns[p.pos].Type == token.INVALID {
		p.printInvalid()
	}
}
