package main

import (
	"fmt"
	"io"
	"os"

	"github.com/claudemuller/oohooh-aahaah-go/internal/pkg/evaluator"
	"github.com/claudemuller/oohooh-aahaah-go/internal/pkg/lexer"
	"github.com/claudemuller/oohooh-aahaah-go/internal/pkg/object"
	"github.com/claudemuller/oohooh-aahaah-go/internal/pkg/parser"
	"github.com/claudemuller/oohooh-aahaah-go/internal/tools/repl"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <file.mon>/n", os.Args[0])
		os.Exit(0)
	}

	code, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}

	env := object.NewEnvironment()
	l := lexer.New(string(code))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		repl.PrintParserErrors(os.Stdout, p.Errors())
		os.Exit(1)
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	}
}
