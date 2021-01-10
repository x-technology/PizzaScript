package repl

import (
	"bufio"
	"fmt"
	"io"
	"pizzascript/lexer"
)

const PROMPT = ">> "
// 1+2

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// debug
	l := lexer.New("1=0")
	l.Print()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		l.Print()
	}
}
