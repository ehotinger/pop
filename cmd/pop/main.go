package main

import (
	"log"

	"github.com/ehotinger/pop/expr"
)

func main() {
	parser, err := expr.NewExpressionParser(`100 > 0`)
	if err != nil {
		log.Fatal(err)
	}

	expr, err := parser.ParseExpression()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("NodeType:", expr.NodeType(), "Type:", expr.Type().ToString(), "--", expr.ToString())

	// for _, token := range parser.GetTokens() {
	// 	log.Printf("token - kind: %s, text: %s", token.Type.ToString(), token.Text)
	// }

	// result, err := parser.Evaluate(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Result:", result)
}
