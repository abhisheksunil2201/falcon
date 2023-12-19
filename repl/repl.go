package repl

import (
	"bufio"
	"falcon/evaluator"
	"falcon/lexer"
	"falcon/object"
	"falcon/parser"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
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
		//parse
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const FALCON = `
.------._ 
.-"""'-.<')    '-._ 
(.--. _   '._       ''---.__.-'
'   ';'-.-'         '-    ._
  .--'''  '._      - '   .
   '""'-.    '---'    ,
		 '\
		   '\      .'
			 ''. '
				 ''.`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, FALCON)
	io.WriteString(out, "Woops! We ran into some falcon feathers here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
