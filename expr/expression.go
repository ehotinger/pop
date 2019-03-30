package expr

import (
	"reflect"
)

type ExpressionType int

const (
	AddExpr ExpressionType = iota
	AddCheckedExpr
	AndExpr
	AndAlsoExpr // Short circuiting
	ConstantExpr
	DivideExpr
	EqualExpr
	ExclusiveOrExpr
	GreaterThanExpr
	GreaterThanOrEqualExpr
	LeftShiftExpr
	LessThanExpr
	LessThanOrEqualExpr
	ModuloExpr
	MultiplyExpr
	MultiplyCheckedExpr
	NegateExpr
	UnaryPlusExpr // TODO: support this?
	NotExpr
	NotEqualExpr
	OrExpr
	OrElseExpr // Short circuiting
	ParameterExpr
	PowerExpr
	RightShiftExpr
	SubtractExpr
	SubtractCheckedExpr
	UnknownExpr
)

type Expression interface {
	ToString() string
	Kind() reflect.Kind
}

type AbstractExpression struct {
	nodeType ExpressionType
	kind     reflect.Kind
}

func (e *AbstractExpression) ToString() string {
	if e == nil {
		return "UnknownExpr"
	}
	switch e.nodeType {
	case AddExpr:
		return "AddExpr"
	case AddCheckedExpr:
		return "AddCheckedExpr"
	case AndExpr:
		return "AndExpr"
	case AndAlsoExpr:
		return "AndAlsoExpr"
	case ConstantExpr:
		return "ConstantExpr"
	case DivideExpr:
		return "DivideExpr"
	case EqualExpr:
		return "EqualExpr"
	case ExclusiveOrExpr:
		return "ExclusiveOrExpr"
	case GreaterThanExpr:
		return "GreaterThanExpr"
	case GreaterThanOrEqualExpr:
		return "GreaterThanOrEqualExpr"
	case LeftShiftExpr:
		return "LeftShiftExpr"
	case LessThanExpr:
		return "LessThanExpr"
	case LessThanOrEqualExpr:
		return "LessThanOrEqualExpr"
	case ModuloExpr:
		return "ModuloExpr"
	case MultiplyExpr:
		return "MultiplyExpr"
	case MultiplyCheckedExpr:
		return "MultiplyCheckedExpr"
	case NegateExpr:
		return "NegateExpr"
	case UnaryPlusExpr:
		return "UnaryPlusExpr"
	case NotExpr:
		return "NotExpr"
	case NotEqualExpr:
		return "NotEqualExpr"
	case OrExpr:
		return "OrExpr"
	case OrElseExpr:
		return "OrElseExpr"
	case ParameterExpr:
		return "ParameterExpr"
	case PowerExpr:
		return "PowerExpr"
	case RightShiftExpr:
		return "RightShiftExpr"
	case SubtractExpr:
		return "SubtractExpr"
	case SubtractCheckedExpr:
		return "SubtractCheckedExpr"
	default:
		return "UnknownExpr"
	}
}

func (e *AbstractExpression) Kind() reflect.Kind {
	return e.kind
}

func IsArithmetic(t reflect.Kind) bool {
	switch t {
	case reflect.Int:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		return true
	default:
		return false
	}
}
