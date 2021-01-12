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
	input := "11=00"
	fmt.Println(input)
	l := lexer.New(input)
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
