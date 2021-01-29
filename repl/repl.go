package repl

import (
	"bufio"
	"fmt"
	"io"
	"pizzascript/lexer"
	"pizzascript/parser"
	"pizzascript/utils/log"
)

const PROMPT = ">> "

// Start executes repl, to try PizzaScript with standard input/output
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// debug
	input := "1 + 2"
	log.Info("lexer & parser example", input)

	l := lexer.New(input)
	l.Print()
	p := parser.New(l)
	p.Print()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		l.Print()

		p := parser.New(l)
		p.Print()
	}
}
