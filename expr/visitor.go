package expr

import (
	"errors"
	"fmt"
	"log"
)

type Visitor interface {
	Visit() (interface{}, error)
}

func CreateVisitorFromExpression(node Expression) (Visitor, error) {
	switch node.Type() {
	case UnknownExpr:
		return nil, errors.New("unable to create visitor for unknown expression")
	case ConstantExpr:
		return NewConstantVisitor(node.(*ConstantExpression)), nil
	case ParameterExpr:
		return NewParameterVisitor(node.(*ParameterExpression)), nil
	default:
		return NewBinaryVisitor(node.(*BinaryExpression)), nil
	}
}

type BinaryVisitor struct {
	root *BinaryExpression
}

func NewBinaryVisitor(root *BinaryExpression) *BinaryVisitor {
	return &BinaryVisitor{
		root: root,
	}
}

func (v *BinaryVisitor) Visit() (interface{}, error) {
	var left Visitor
	left, err := CreateVisitorFromExpression(v.root.left)
	if err != nil {
		return nil, err
	}
	lVal, lErr := left.Visit()
	if lErr != nil {
		return nil, lErr
	}
	var right Visitor
	right, err = CreateVisitorFromExpression(v.root.right)
	if err != nil {
		return nil, err
	}
	rVal, rErr := right.Visit()
	if rErr != nil {
		return nil, rErr
	}
	log.Println(v.root.String())

	switch v.root.Type() {
	case AddExpr:
		return lVal.(int) + rVal.(int), nil
	case AddCheckedExpr:
		return nil, errors.New("unimplemented")
	case AndAlsoExpr:
		return nil, errors.New("unimplemented")
	case DivideExpr:
		return lVal.(int) / rVal.(int), nil
	case EqualExpr:
		return true, errors.New("unimplemented")
	case ExclusiveOrExpr:
		return nil, errors.New("unimplemented")
	case GreaterThanExpr:
		return true, errors.New("unimplemented")
	case GreaterThanOrEqualExpr:
		return true, errors.New("unimplemented")
	case LessThanExpr:
		return true, errors.New("unimplemented")
	case LessThanOrEqualExpr:
		return true, errors.New("unimplemented")
	case MultiplyExpr:
		return lVal.(int) * rVal.(int), nil
	case MultiplyCheckedExpr:
		return nil, errors.New("unimplemented")
	case OrExpr:
		return true, errors.New("unimplemented")
	case OrElseExpr:
		return true, errors.New("unimplemented")
	case PowerExpr:
		return nil, errors.New("unimplemented")
	case SubtractExpr:
		return lVal.(int) - rVal.(int), nil
	case SubtractCheckedExpr:
		return nil, errors.New("unimplemented")
	}

	return nil, fmt.Errorf("unknown expression type: %v", v.root.Type())
}

type ConstantVisitor struct {
	root *ConstantExpression
}

func NewConstantVisitor(root *ConstantExpression) *ConstantVisitor {
	return &ConstantVisitor{
		root: root,
	}
}

func (v *ConstantVisitor) Visit() (interface{}, error) {
	log.Println("[constant]", v.root)
	return v.root.value, nil
}

type ParameterVisitor struct {
	root *ParameterExpression
}

func NewParameterVisitor(root *ParameterExpression) *ParameterVisitor {
	return &ParameterVisitor{
		root: root,
	}
}

func (v *ParameterVisitor) Visit() (interface{}, error) {
	log.Println("[param]", v.root)
	return false, nil
}
