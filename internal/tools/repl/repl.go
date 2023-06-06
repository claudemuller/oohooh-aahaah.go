package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/claudemuller/oohooh-aahaah-go/internal/pkg/lexer"
	"github.com/claudemuller/oohooh-aahaah-go/internal/pkg/token"
)

func Start(in io.Reader, out io.Writer) {
	const PROMPT = ">> "

	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
