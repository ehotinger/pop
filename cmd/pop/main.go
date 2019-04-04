package main

import (
	"log"

	"github.com/ehotinger/pop/expr"
)

func main() {
	parameters := make(map[string]interface{})
	parameters["a"] = 3
	parameters["b"] = 4
	parser, err := expr.NewExpressionParser(`a == b`, parameters)
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
	if err := visitor.Visit(); err != nil {
		log.Fatalf("failed traversal: %v", err)
	}

	// for _, token := range parser.GetTokens() {
	// 	log.Printf("token - kind: %s, text: %s", token.Type, token.Text)
	// }

	// result, err := parser.Evaluate(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Result:", result)
}
