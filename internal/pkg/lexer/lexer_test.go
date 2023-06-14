package lexer

import (
	"testing"

	"github.com/claudemuller/oohooh-aahaah-go/internal/pkg/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
[1, 2];
{"foo": "bar"}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LetTok, "let"},
		{token.IdentTok, "five"},
		{token.AssignTok, "="},
		{token.IntTok, "5"},
		{token.SemicolonTok, ";"},
		{token.LetTok, "let"},
		{token.IdentTok, "ten"},
		{token.AssignTok, "="},
		{token.IntTok, "10"},
		{token.SemicolonTok, ";"},
		{token.LetTok, "let"},
		{token.IdentTok, "add"},
		{token.AssignTok, "="},
		{token.FunctionTok, "fn"},
		{token.LParenTok, "("},
		{token.IdentTok, "x"},
		{token.CommaTok, ","},
		{token.IdentTok, "y"},
		{token.RParenTok, ")"},
		{token.LBraceTok, "{"},
		{token.IdentTok, "x"},
		{token.PlusTok, "+"},
		{token.IdentTok, "y"},
		{token.SemicolonTok, ";"},
		{token.RBraceTok, "}"},
		{token.SemicolonTok, ";"},
		{token.LetTok, "let"},
		{token.IdentTok, "result"},
		{token.AssignTok, "="},
		{token.IdentTok, "add"},
		{token.LParenTok, "("},
		{token.IdentTok, "five"},
		{token.CommaTok, ","},
		{token.IdentTok, "ten"},
		{token.RParenTok, ")"},
		{token.SemicolonTok, ";"},
		{token.BangTok, "!"},
		{token.MinusTok, "-"},
		{token.SlashTok, "/"},
		{token.AsteriskTok, "*"},
		{token.IntTok, "5"},
		{token.SemicolonTok, ";"},
		{token.IntTok, "5"},
		{token.LTTok, "<"},
		{token.IntTok, "10"},
		{token.GTTok, ">"},
		{token.IntTok, "5"},
		{token.SemicolonTok, ";"},
		{token.IfTok, "if"},
		{token.LParenTok, "("},
		{token.IntTok, "5"},
		{token.LTTok, "<"},
		{token.IntTok, "10"},
		{token.RParenTok, ")"},
		{token.LBraceTok, "{"},
		{token.ReturnTok, "return"},
		{token.TrueTok, "true"},
		{token.SemicolonTok, ";"},
		{token.RBraceTok, "}"},
		{token.ElseTok, "else"},
		{token.LBraceTok, "{"},
		{token.ReturnTok, "return"},
		{token.FalseTok, "false"},
		{token.SemicolonTok, ";"},
		{token.RBraceTok, "}"},
		{token.IntTok, "10"},
		{token.EQTok, "=="},
		{token.IntTok, "10"},
		{token.SemicolonTok, ";"},
		{token.IntTok, "10"},
		{token.NotEQTok, "!="},
		{token.IntTok, "9"},
		{token.SemicolonTok, ";"},
		{token.StringTok, "foobar"},
		{token.StringTok, "foo bar"},
		{token.LBracketTok, "["},
		{token.IntTok, "1"},
		{token.CommaTok, ","},
		{token.IntTok, "2"},
		{token.RBracketTok, "]"},
		{token.SemicolonTok, ";"},
		{token.LBraceTok, "{"},
		{token.StringTok, "foo"},
		{token.ColonTok, ":"},
		{token.StringTok, "bar"},
		{token.RBraceTok, "}"},
		{token.EOFTok, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
