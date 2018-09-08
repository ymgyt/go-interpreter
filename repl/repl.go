package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ymgyt/go-interpreter/evaluator"
	"github.com/ymgyt/go-interpreter/lexer"
	"github.com/ymgyt/go-interpreter/object"
	"github.com/ymgyt/go-interpreter/parser"
)

const PROMPT = ">> "

func Start(r io.Reader, w io.Writer) {
	s := bufio.NewScanner(r)
	env := object.NewEnvironment()

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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(w, evaluated.Inspect())
			io.WriteString(w, "\n")
		} else {
			io.WriteString(w, "nil\n")
		}

	}
}

func printParserErrors(w io.Writer, errors []error) {
	for _, err := range errors {
		io.WriteString(w, "\t"+err.Error()+"\n")
	}
}
