package main

import (
	"log"

	"github.com/ehotinger/pop/expr"
)

func main() {
	parameters := make(map[string]interface{})
	parameters["a"] = 10
	parameters["b"] = 2
	parser, err := expr.NewExpressionParser(`a + b * 30 + 5 - 20 - 900`, parameters)
	if err != nil {
		log.Fatal(err)
	}

	expression, err := parser.ParseExpression()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("NodeType:", expression.NodeType(), "Type:", expression.Type(), "--", expression)

	visitor, err := expr.CreateVisitorFromExpression(expression)
	if err != nil {
		log.Fatalf("failed to create visitor: %v", err)
	}
	val, err := visitor.Visit()
	if err != nil {
		log.Fatalf("failed traversal: %v", err)
	}

	log.Println("Interpreted value: ", val)

	// for _, token := range parser.GetTokens() {
	// 	log.Printf("token - kind: %s, text: %s", token.Type, token.Text)
	// }

	// result, err := parser.Evaluate(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Result:", result)
}
