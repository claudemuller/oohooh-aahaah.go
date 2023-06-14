package ast

import (
	"testing"

	"github.com/claudemuller/oohooh-aahaah-go/internal/pkg/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LetTok, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IdentTok, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IdentTok, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
