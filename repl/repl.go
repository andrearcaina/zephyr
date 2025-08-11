package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/andrearcaina/zephyr/lexer"
	"github.com/andrearcaina/zephyr/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// for each token in the input line, such that the token type is not EOF, print that token
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
