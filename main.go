package main

import (
	"fmt"

	"5100/lexis"
	"5100/lexis/lexstates"
	"5100/semantics"
	"5100/syntax"
)

func main() {
	var err error
	for input := ""; input != "exit"; {
		fmt.Println()
		fmt.Print("<< ")

		_, err = fmt.Scanln(&input)
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
