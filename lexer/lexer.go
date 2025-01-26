package lexer

import (
	"eg-lsp/token"
	"log"
	"strconv"
	"unicode"
)

func addNewLine(curr *string, tkns *[]token.Token) {
	token.Add(curr, tkns)
	token.NewLine(tkns)
}

func addComment(curr, input *string, tkns *[]token.Token, i *int, n int) {
	token.Add(curr, tkns)
	for (*input)[*i] != '\n' && (*i) < n {
		(*i)++
	}
	(*tkns) = append(*tkns, token.Create("//"))
	token.NewLine(tkns)
}

func lexString(curr, input *string, tkns *[]token.Token, i *int, n int) {
	token.Add(curr, tkns)
	j := *i
	(*i)++
	for (*input)[*i] != '\n' && string((*input)[*i]) != `"` && (*i) < n {
		(*i)++
	}
	if string((*input)[*i]) == `"` {
		(*tkns) = append(*tkns, token.String((*input)[j:*i+1]))
		return
	}
	(*tkns) = append(*tkns, token.Invalid((*input)[j:*i+1]))
}

// checks and analyzes two character operators like ==, !=, etc.
func twoCharOpts(ch string, input *string, tkns *[]token.Token, i *int, n int) bool {
	if (*i)+1 < n {
		nextCh := string((*input)[(*i)+1])
		twoChar := ch + nextCh
		if twoChar == "==" || twoChar == "!=" || twoChar == "<=" || twoChar == ">=" {
			*tkns = append(*tkns, token.Create(twoChar))
			(*i)++
			return true
		}
	}
	return false
}

func canBuildNegative(tkLen int, prevTknType string) bool {
	return tkLen == 0 || prevTknType == token.EQ || prevTknType == token.LPAREN ||
		prevTknType == token.COMMA || prevTknType == token.RETURN
}

func lexDash(curr, input *string, tkns *[]token.Token, i *int, n int) {
	token.Add(curr, tkns)
	// Check if this might be a negative number
	tkLen := len(*tkns)
	prevTknType := (*tkns)[tkLen-1].Type
	var nextTknLiteral rune
	if *i+1 < n {
		nextTknLiteral = rune((*input)[*i+1])
	}
	if IsDigit(nextTknLiteral) {
		if canBuildNegative(tkLen, prevTknType) {
			*curr = "-" // Start building a negative number
			return
		}
	}
	*tkns = append(*tkns, token.Create("-")) // It's an operator
}

func isComment(input *string, i *int, n int) bool {
	return *i+1 < n-1 && string((*input)[*i+1]) == "/"
}

func analyzeAndAct(curr, input *string, tkns *[]token.Token, i *int, n int) {
	if (*input)[*i] == '\n' {
		addNewLine(curr, tkns)
		return
	}
	switch string((*input)[*i]) {
	case " ", "\t":
		token.Add(curr, tkns)
	case "(", ")", "+", "*", "=", "<", ">", "!", ",", "[", "]", "{", "}", ":":
		token.Add(curr, tkns)
		exists := twoCharOpts(string((*input)[*i]), input, tkns, i, n)
		if !exists {
			*tkns = append(*tkns, token.Create(string((*input)[*i])))
		}
	case "/":
		if isComment(input, i, n) {
			addComment(curr, input, tkns, i, n)
			return
		} // Else it is a divide sign.
		token.Add(curr, tkns)
		*tkns = append(*tkns, token.Create("/"))
	case `"`:
		lexString(curr, input, tkns, i, n)
	case "-":
		lexDash(curr, input, tkns, i, n)
	// TODO: case ".": HANDLE dot cases once methods are implemented.
	default:
		*curr += string((*input)[*i])
	}
}

func Tokenize(input string) []token.Token {
	if len(input) == 0 {
		return token.EndOfFile()
	}
	tkns := []token.Token{}
	curr := ""
	for i := 0; i < len(input); i++ {
		analyzeAndAct(&curr, &input, &tkns, &i, len(input))
	}
	return tkns
}

func ConvToFloat(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("Failed to convert %v to float", s)
	}
	return n
}

func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}
