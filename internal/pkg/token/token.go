package token

// Might not want to hide the type here. And probably more
// performant to use an int or byte.
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	IllegalTok = "ILLEGAL"
	EOFTok     = "EOF"
	// Identifiers + literals.
	IdentTok = "IDENT" // e.g. add, getUsers, x, y
	IntTok   = "INT"   // 1234
	// Operators.
	AssignTok   = "="
	PlusTok     = "+"
	MinusTok    = "-"
	BangTok     = "!"
	AsteriskTok = "*"
	SlashTok    = "/"
	LTTok       = "<"
	GTTok       = ">"
	EQTok       = "=="
	NotEQTok    = "!="
	// Delimiters.
	CommaTok     = ","
	SemicolonTok = ";"
	LParenTok    = "("
	RParenTok    = ")"
	LBraceTok    = "{"
	RBraceTok    = "}"
	LBracketTok  = "["
	RBracketTok  = "]"
	ColonTok     = ":"
	// Language keywords.
	FunctionTok = "FUNCTION"
	LetTok      = "LET"
	TrueTok     = "TRUE"
	FalseTok    = "FALSE"
	IfTok       = "IF"
	ElseTok     = "ELSE"
	ReturnTok   = "RETURN"
	StringTok   = "STRING"
)

func LookupIdent(ident string) TokenType {
	var keywords = map[string]TokenType{
		"fn":     FunctionTok,
		"let":    LetTok,
		"true":   TrueTok,
		"false":  FalseTok,
		"if":     IfTok,
		"else":   ElseTok,
		"return": ReturnTok,
	}

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IdentTok
}
