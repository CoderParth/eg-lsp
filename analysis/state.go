package analysis

import (
	"eg-lsp/lexer"
	"eg-lsp/lsp"
	"eg-lsp/parser"
)

type State struct {
	// Map of file names to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	tkns := lexer.Tokenize(text)
	p := parser.NewParser(tkns)
	p.Parse()
	for lineNo, msg := range p.Errors {
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Range:    LineRange(lineNo-1, 0, 5),
			Severity: 1,
			Source:   "EG-LSP",
			Message:  msg,
		})
	}
	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}
