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
	case NegateExpr:
		fallthrough
	case UnaryPlusExpr:
		fallthrough
	case NotExpr:
		return NewUnaryVisitor(node.(*UnaryExpression)), nil
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
		lInt, rInt, err := convertExpressionToInt(lVal, rVal)
		if err != nil {
			return nil, err
		}
		return lInt + rInt, nil
	case AddCheckedExpr:
		return nil, errors.New("unimplemented")
	case AndAlsoExpr:
		return nil, errors.New("unimplemented")
	case DivideExpr:
		lInt, rInt, err := convertExpressionToInt(lVal, rVal)
		if err != nil {
			return nil, err
		}
		return lInt / rInt, nil
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
		lInt, rInt, err := convertExpressionToInt(lVal, rVal)
		if err != nil {
			return nil, err
		}
		return lInt * rInt, nil
	case MultiplyCheckedExpr:
		return nil, errors.New("unimplemented")
	case OrExpr:
		return true, errors.New("unimplemented")
	case OrElseExpr:
		return true, errors.New("unimplemented")
	case PowerExpr:
		return nil, errors.New("unimplemented")
	case SubtractExpr:
		lInt, rInt, err := convertExpressionToInt(lVal, rVal)
		if err != nil {
			return nil, err
		}
		return lInt - rInt, nil
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
	return false, nil
}

type UnaryVisitor struct {
	root *UnaryExpression
}

func NewUnaryVisitor(root *UnaryExpression) *UnaryVisitor {
	return &UnaryVisitor{
		root: root,
	}
}

func (v *UnaryVisitor) Visit() (interface{}, error) {
	switch v.root.Type() {
	case NegateExpr:
		return nil, errors.New("unimplemented")
	case UnaryPlusExpr:
		return nil, errors.New("unimplemented")
	case NotExpr:
		return nil, errors.New("unimplemented")
	}

	return nil, fmt.Errorf("unknown expression type: %v", v.root.Type())
}

func convertToInt(val interface{}) (int, error) {
	switch t := val.(type) {
	case int:
		return t, nil
	case uint:
		return int(t), nil
	case int8:
		return int(t), nil
	case uint8:
		return int(t), nil
	case int16:
		return int(t), nil
	case uint16:
		return int(t), nil
	case int32:
		return int(t), nil
	case uint32:
		return int(t), nil
	case int64:
		return int(t), nil
	case uint64:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	}

	return 0, fmt.Errorf("unable to convert value to integer: %v", val)
}

func convertExpressionToInt(leftVal interface{}, rightVal interface{}) (left int, right int, err error) {
	left, err = convertToInt(leftVal)
	if err != nil {
		return left, right, err
	}
	right, err = convertToInt(rightVal)
	return left, right, err
}
