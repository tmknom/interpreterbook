package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			return
		}

		l := lexer.NewLexer(line)
		for tok := l.NextToken(); !tok.IsEOF(); tok = l.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", tok)
			if err != nil {
				panic(err)
			}
		}
	}
}
