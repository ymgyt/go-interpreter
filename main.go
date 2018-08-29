package main

import (
	"fmt"
	"os"

	"github.com/ymgyt/go-interpreter/repl"
)

func main() {
	fmt.Printf("type input source code\n")
	repl.Start(os.Stdin, os.Stdout)
}
