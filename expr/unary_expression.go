package expr

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	errInvalidExpression = errors.New("expression can't be nil")
)

type UnaryExpression struct {
	self    *AbstractExpression
	operand Expression
}

func NewUnaryExpression(operand Expression, nodeType ExpressionType, kind reflect.Kind) *UnaryExpression {
	return &UnaryExpression{
		self: &AbstractExpression{
			nodeType: nodeType,
			kind:     kind,
		},
		operand: operand,
	}
}

func (e *UnaryExpression) Kind() reflect.Kind {
	return e.self.kind
}

func (e *UnaryExpression) Type() ExpressionType {
	return e.self.nodeType
}

func (e *UnaryExpression) ToString() string {
	if e == nil {
		return "<nil>"
	}
	switch e.self.nodeType {
	case NotExpr:
		return fmt.Sprintf("Not(%s)", e.operand.ToString())
	case NegateExpr:
		return fmt.Sprintf("-%s", e.operand.ToString())
	case UnaryPlusExpr:
		return fmt.Sprintf("+%s", e.operand.ToString())
	default:
		return fmt.Sprintf("(%s)", e.operand.ToString())
	}
}

func CreateNegate(expr Expression) (Expression, error) {
	if expr == nil {
		return nil, errInvalidExpression
	}
	return NewUnaryExpression(expr, NegateExpr, expr.Kind()), nil
}

func CreateNot(expr Expression) (Expression, error) {
	if expr == nil {
		return nil, errInvalidExpression
	}
	return NewUnaryExpression(expr, NotExpr, expr.Kind()), nil
}
