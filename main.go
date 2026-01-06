package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"5100/lexis"
	"5100/lexis/lexstates"
	"5100/semantics"
	"5100/syntax"
)

func main() {
	var err error
	rdr := bufio.NewReader(os.Stdin)
	for input := ""; input != "exit"; {
		fmt.Println()
		fmt.Print("<< ")

		input, err = rdr.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		lexer := lexis.NewLexer(input, *lexstates.NewValueState())

		tokens, err := lexer.Parse()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		expr := syntax.Parse(tokens)
		result, err := semantics.Evaluate(expr)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(">>", result)
	}
}
