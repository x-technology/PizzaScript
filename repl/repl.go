package repl

import (
	"bufio"
	"fmt"
	"io"
	"pizzascript/eval"
	"pizzascript/lexer"
	"pizzascript/parser"
)

const PROMPT = "\n>> "

// Start executes repl, to try PizzaScript with standard input/output
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		// Compile(line)
		l := lexer.New(line)

		p := parser.New(l)
		// fmt.Println(p.Print())

		e := eval.New(p)

		fmt.Println(e.Eval())
	}
}

// Compile takes a string, and compiles to WebAssembly, returns an output in wat format
// TODO move to other module, should not be in repl
func Compile(input string) string {
	l := lexer.New(input)

	p := parser.New(l)
	return p.PrintWat()
}
