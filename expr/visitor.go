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
		return NewConstantVisitor(node.(*ConstantExpression)), nil
	case ParameterExpr:
		return NewParameterVisitor(node.(*ParameterExpression)), nil
	case AddExpr:
		return NewBinaryVisitor(node.(*BinaryExpression)), nil
	}

	return nil, fmt.Errorf("unable to create visitor for %v", node.Type())
}

type BinaryVisitor struct {
	root *BinaryExpression
}

func NewBinaryVisitor(root *BinaryExpression) *BinaryVisitor {
	return &BinaryVisitor{
		root: root,
	}
}

func (v *BinaryVisitor) Visit() (err error) {
	log.Println("[bin]")
	var left Visitor
	left, err = CreateVisitorFromExpression(v.root.left)
	if err != nil {
		return err
	}
	log.Println("[left]")
	err = left.Visit()
	if err != nil {
		return err
	}
	var right Visitor
	right, err = CreateVisitorFromExpression(v.root.right)
	if err != nil {
		return err
	}
	log.Println("[right]")
	return right.Visit()
}

type ConstantVisitor struct {
	root *ConstantExpression
}

func NewConstantVisitor(root *ConstantExpression) *ConstantVisitor {
	return &ConstantVisitor{
		root: root,
	}
}

func (v *ConstantVisitor) Visit() (err error) {
	log.Println("[constant]", v.root)
	return nil
}

type ParameterVisitor struct {
	root *ParameterExpression
}

func NewParameterVisitor(root *ParameterExpression) *ParameterVisitor {
	return &ParameterVisitor{
		root: root,
	}
}

func (v *ParameterVisitor) Visit() (err error) {
	log.Println("[param]", v.root)
	return nil
}
