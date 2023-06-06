package token

// Might not want to hide the type here. And probably more
// performant to use an int or byte.
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals.
	IDENT = "IDENT" // e.g. add, getUsers, x, y
	INT   = "INT"   // 1234
	// Operators.
	ASSIGN = "="
	PLUS   = "+"
	// Delimiters.
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Language keywords.
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

func LookupIdent(ident string) TokenType {
	var keywords = map[string]TokenType{
		"fn":  FUNCTION,
		"let": LET,
	}

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
