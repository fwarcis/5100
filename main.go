package main

import (
	"fmt"
	"log"

	"5100/lexis"
	"5100/semantics"
)

func main() {
	var err error
	for input := ""; input != "exit"; {
		fmt.Print("<< ")

		_, err = fmt.Scanln(&input)
		if err != nil {
			log.Println(err.Error())
		}

		lexer := lexis.NewLexer(input, *lexis.NewValueState())

		tokens, err := lexer.Parse()
		if err != nil {
			fmt.Println(err.Error())
		}
		expr := semantics.Parse(tokens)
		result := expr.Evaluate()

		fmt.Println(">>", result)
		fmt.Println()
	}
}
