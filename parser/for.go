package parser

import (
	"eg-lsp/token"
	"fmt"
)

func (p *Parser) parseFor() {
	errorLine := p.lineNo
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Unexpected end of for loop"
		return
	}
	p.pos++ // Skip FOR
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Incomplete for loop condition"
		return
	}
	p.parseComparison()
	p.pos++
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Expected 'do' keyword"
		return
	}
	if p.tkns[p.pos].Type != token.DO {
		p.Errors[errorLine] = fmt.Sprintf("Expected 'do', got '%s'", p.tkns[p.pos].Literal)
		return
	}
	p.pos++ // Skip DO
	p.parseForBlock()
}

func (p *Parser) parseForBlock() {
	errorLine := p.lineNo
	if p.pos >= len(p.tkns) {
		p.Errors[errorLine] = "Unexpected end of for block"
		return
	}
	// Track start of block
	blockStart := p.pos
	p.parseBlock()
	// Verify we found END
	if p.pos >= len(p.tkns) {
		p.Errors[blockStart] = "For loop missing 'end' keyword"
		return
	}
	if p.tkns[p.pos].Type != token.END {
		p.Errors[blockStart] = "For loop missing 'end' keyword"
		return
	}
	p.pos++ // Skip END
	if p.pos < len(p.tkns) && p.tkns[p.pos].Type == token.NEWLINE {
		p.lineNo++
		p.pos++
	}
}
