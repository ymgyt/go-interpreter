package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ymgyt/go-interpreter/lexer"
	"github.com/ymgyt/go-interpreter/token"
)

const PROMPT = ">> "

func Start(r io.Reader, w io.Writer) {
	s := bufio.NewScanner(r)

	for {
		fmt.Fprintf(w, PROMPT)
		scanned := s.Scan()
		if !scanned {
			return
		}

		line := s.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(w, "%+v\n", tok)
		}
	}
}
