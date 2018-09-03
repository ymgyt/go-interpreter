package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ymgyt/go-interpreter/lexer"
	"github.com/ymgyt/go-interpreter/parser"
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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(w, p.Errors())
			continue
		}

		io.WriteString(w, program.String())
		io.WriteString(w, "\n")
	}
}

func printParserErrors(w io.Writer, errors []error) {
	for _, err := range errors {
		io.WriteString(w, "\t"+err.Error()+"\n")
	}
}
