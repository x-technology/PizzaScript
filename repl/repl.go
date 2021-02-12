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
		l := lexer.New(line)

		p := parser.New(l)
		e := eval.New(p)

		fmt.Println(e.Eval())
	}
}
