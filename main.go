package main

import (
	"fmt"

	"5100/lexis"
	"5100/semantics"
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

		lexer := lexis.NewLexer(input, *lexis.NewValueState())

		tokens, err := lexer.Parse()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		expr := semantics.Parse(tokens)
		result := expr.Evaluate()

		fmt.Println(">>", result)
	}
}
