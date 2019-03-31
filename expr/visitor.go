package expr

import (
	"fmt"
	"log"
)

type Visitor interface {
	Visit() error
}

func CreateVisitorFromExpression(node Expression) (Visitor, error) {
	switch node.Type() {
	case ConstantExpr:
		log.Println("TODO: ConstantVisitor")
		break
	case ParameterExpr:
		log.Println("TODO: ParameterVisitor")
		break
	case AddExpr:
		return NewBinaryVisitor(node.(*BinaryExpression)), nil
	}

	return nil, fmt.Errorf("unable to create visitor for %v", node.Type())
}
