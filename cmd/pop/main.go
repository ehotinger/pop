package main

import (
	"log"

	"github.com/ehotinger/pop/expr"
)

func main() {
	expression, err := expr.NewExpressionParser(`"100" > 0`)
	if err != nil {
		log.Fatal(err)
	}

	for _, token := range expression.GetTokens() {
		log.Printf("token - kind: %s, text: %s", token.Type.ToString(), token.Text)
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Result:", result)
}
