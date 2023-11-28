package repl

import (
	"bufio"
	"falcon/lexer"
	"falcon/token"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		//print prompt
		fmt.Printf("%s", PROMPT)
		//read input
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		//get input
		line := scanner.Text()
		//create lexer
		l := lexer.New(line)
		//print tokens
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}