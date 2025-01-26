package token

import (
	"strconv"
	"unicode"
)

type Token struct {
	Type    string
	Literal string
}

const (
	EMPTY = "EMPTY"
	EOF   = "EOF"

	NEWLINE = "NEWLINE"
	NUM     = "NUM"
	IDENT   = "IDENT"
	STRING  = "STRING"
	INVALID = "INVALID"
	COMMENT = "COMMENT"
	// Keywords
	PRINT    = "PRINT"
	INPUT    = "INPUT"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	THEN     = "THEN"
	FN       = "FN"
	END      = "END"
	FOR      = "FOR"
	DO       = "DO"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	RETURN   = "RETURN"
	LEN      = "LEN"
	APPEND   = "APPEND"
	POP      = "POP"
	DEL      = "DEL"
	// Operators
	EQ       = "EQ"       // Equals
	PLUS     = "PLUS"     // Plus (+)
	MINUS    = "MINUS"    // Minus (-)
	ASTERISK = "ASTERISK" // Asterisk (*)
	SLASH    = "SLASH"    // Slash (/)
	EQEQ     = "EQEQ"     // Double equals (==)
	NOTEQ    = "NOTEQ"    // Not equals (!=)
	LT       = "LT"       // Less than (<)
	LTEQ     = "LTEQ"     // Less than or equals (<=)
	GT       = "GT"       // Greater than (>)
	GTEQ     = "GTEQ"     // Greater than or equals (>=)
	LPAREN   = "LPAREN"   // Left Parenthesis `(`
	RPAREN   = "RPAREN"   // Right Parenthesis `)`
	LCURLY   = "LCURLY"   // Left curly braces `{`
	RCURLY   = "RCURLY"   // Right curly braches `}`
	COMMA    = "COMMA"
	LBRACK   = "LBRACE" // [
	RBRACK   = "RBRACK" // ]
	DOT      = "DOT"
	COLON    = "COLON"
	AND      = "AND"
	OR       = "OR"
	// TYPE
	NIL   = "NIL"
	TRUE  = "TRUE"
	FALSE = "FALSE"
)

func tokenType(s string) string {
	switch s {
	case "//":
		return COMMENT
	// Operators
	case "=":
		return EQ
	case "+":
		return PLUS
	case "-":
		return MINUS
	case "*":
		return ASTERISK
	case "/":
		return SLASH
	case "==":
		return EQEQ
	case "!=":
		return NOTEQ
	case "<":
		return LT
	case "<=":
		return LTEQ
	case ">":
		return GT
	case ">=":
		return GTEQ
	case "and":
		return AND
	case "or":
		return OR
	case "(":
		return LPAREN
	case ")":
		return RPAREN
	case "[":
		return LBRACK
	case "]":
		return RBRACK
	case ",":
		return COMMA
	case ".":
		return DOT
	case "{":
		return LCURLY
	case "}":
		return RCURLY
	case ":":
		return COLON
	// Keywords
	case "print":
		return PRINT
	case "input":
		return INPUT
	case "fn":
		return FN
	case "end":
		return END
	case "for":
		return FOR
	case "do":
		return DO
	case "if":
		return IF
	case "then":
		return THEN
	case "else":
		return ELSE
	case "break":
		return BREAK
	case "continue":
		return CONTINUE
	case "return":
		return RETURN
	case "len":
		return LEN
	case "append":
		return APPEND
	case "pop":
		return POP
	case "del":
		return DEL
	case "nil":
		return NIL
	case "true":
		return TRUE
	case "false":
		return FALSE
	}
	if isIdent(s) {
		return IDENT
	}
	if isNum(s) {
		return NUM
	}
	return INVALID
}

func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isIdent(s string) bool {
	firstChar := s[0]
	if !unicode.IsLetter(rune(firstChar)) && firstChar != '_' {
		return false
	}
	return true
}

// ---------- Exported Functions ------------------- //
func Create(s string) Token {
	return Token{
		Type:    tokenType(s),
		Literal: s,
	}
}

func String(s string) Token {
	return Token{
		Type:    STRING,
		Literal: s,
	}
}

func Invalid(s string) Token {
	return Token{
		Type:    INVALID,
		Literal: s,
	}
}

func NewLine(tkns *[]Token) {
	*tkns = append(*tkns, Token{Type: NEWLINE, Literal: "\n"})
}

func EndOfFile() []Token {
	tkns := []Token{}
	tkns = append(tkns, Token{Type: EOF, Literal: ""})
	return tkns
}

func Add(curr *string, tkns *[]Token) {
	if *curr != "" {
		*tkns = append(*tkns, Create(*curr))
		*curr = ""
	}
}
