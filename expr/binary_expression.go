package expr

import (
	"fmt"
	"reflect"
)

type BinaryExpression struct {
	left  Expression
	right Expression
	self  *AbstractExpression
}

func NewBinaryExpression(nodeType ExpressionType, left Expression, right Expression, kind reflect.Kind) *BinaryExpression {
	return &BinaryExpression{
		left:  left,
		right: right,
		self: &AbstractExpression{
			nodeType: nodeType,
			Kind:     kind,
		},
	}
}

func (e *BinaryExpression) Left() Expression {
	return e.left
}

func (e *BinaryExpression) Right() Expression {
	return e.right
}

func (e *BinaryExpression) ToString() string {
	operator := e.GetOperator()
	if operator == "" {
		return fmt.Sprintf("%s (%s, %s)", e.self.ToString(), e.left.ToString(), e.right.ToString())
	}
	return fmt.Sprintf("(%s %s %s)", e.left.ToString(), operator, e.right.ToString())
}

func (e *BinaryExpression) GetOperator() string {
	switch e.self.nodeType {
	case AddExpr:
		fallthrough
	case AddCheckedExpr:
		return "+"
	case SubtractExpr:
		fallthrough
	case SubtractCheckedExpr:
		return "-"
	case MultiplyExpr:
		fallthrough
	case MultiplyCheckedExpr:
		return "*"
	case DivideExpr:
		return "/"
	case ModuloExpr:
		return "%"
	case ExclusiveOrExpr:
		fallthrough
	case PowerExpr:
		return "^"
	case AndExpr:
		return "&" // TODO
	case AndAlsoExpr:
		return "&&"
	case OrExpr:
		return "|" // TODO
	case OrElseExpr:
		return "||"
	case LessThanExpr:
		return "<"
	case LessThanOrEqualExpr:
		return "<="
	case GreaterThanExpr:
		return ">"
	case GreaterThanOrEqualExpr:
		return ">="
	case EqualExpr:
		return "="
	case NotEqualExpr:
		return "!="
	case LeftShiftExpr:
		return "<<"
	case RightShiftExpr:
		return ">>"
	default:
		return ""
	}
}
