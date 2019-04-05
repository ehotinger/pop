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

func (e *UnaryExpression) NodeType() string {
	return "UnaryExpression"
}

func (e *UnaryExpression) String() string {
	if e == nil {
		return "<nil>"
	}
	switch e.self.nodeType {
	case NotExpr:
		return fmt.Sprintf("Not(%v)", e.operand)
	case NegateExpr:
		return fmt.Sprintf("-%v", e.operand)
	case UnaryPlusExpr:
		return fmt.Sprintf("+%v", e.operand)
	default:
		return fmt.Sprintf("unary(%v)", e.operand)
	}
}

func CreateUnaryPlus(expr Expression) (Expression, error) {
	if expr == nil {
		return nil, errInvalidExpression
	}
	if IsArithmetic(expr.Kind()) {
		return NewUnaryExpression(expr, UnaryPlusExpr, expr.Kind()), nil
	}
	return nil, errors.New("unary plus not supported for non-arithmetic values")
}

func CreateUnaryNegate(expr Expression) (Expression, error) {
	if expr == nil {
		return nil, errInvalidExpression
	}
	if IsArithmetic(expr.Kind()) && IsUnsigned(expr.Kind()) {
		return NewUnaryExpression(expr, NegateExpr, expr.Kind()), nil
	}
	// TODO: get user defined operator or throw
	return nil, errors.New("unary negate not supported for non-arithmetic or unsigned values")
}

func CreateUnaryNot(expr Expression) (Expression, error) {
	if expr == nil {
		return nil, errInvalidExpression
	}
	return NewUnaryExpression(expr, NotExpr, expr.Kind()), nil
}
