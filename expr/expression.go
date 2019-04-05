package expr

import (
	"reflect"
)

type ExpressionType int

const (
	AddExpr ExpressionType = iota
	AddCheckedExpr
	AndExpr
	AndAlsoExpr // TODO: Short circuiting
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
	OrElseExpr // TODO: Short circuiting
	ParameterExpr
	PowerExpr
	RightShiftExpr
	SubtractExpr
	SubtractCheckedExpr
	UnknownExpr
)

const (
	AddExprString                = "AddExpr"
	AddCheckedExprString         = "AddCheckedExpr"
	AndExprString                = "AndExpr"
	AndAlsoExprString            = "AndAlsoExpr"
	ConstantExprString           = "ConstantExpr"
	DivideExprString             = "DivideExpr"
	EqualExprString              = "EqualExpr"
	ExclusiveOrExprString        = "ExclusiveOrExpr"
	GreaterThanExprString        = "GreaterThanExpr"
	GreaterThanOrEqualExprString = "GreaterThanOrEqualExpr"
	LeftShiftExprString          = "LeftShiftExpr"
	LessThanExprString           = "LessThanExpr"
	LessThanOrEqualExprString    = "LessThanOrEqualExpr"
	ModuloExprString             = "ModuloExpr"
	MultiplyExprString           = "MultiplyExpr"
	MultiplyCheckedExprString    = "MultiplyCheckedExpr"
	NegateExprString             = "NegateExpr"
	UnaryPlusExprString          = "UnaryPlusExpr"
	NotExprString                = "NotExpr"
	NotEqualExprString           = "NotEqualExpr"
	OrExprString                 = "OrExpr"
	OrElseExprString             = "OrElseExpr"
	ParameterExprString          = "ParameterExpr"
	PowerExprString              = "PowerExpr"
	RightShiftExprString         = "RightShiftExpr"
	SubtractExprString           = "SubtractExpr"
	SubtractCheckedExprString    = "SubtractCheckedExpr"
	UnknownExprString            = "UnknownExpr"
)

func (t ExpressionType) String() string {
	switch t {
	case AddExpr:
		return AddExprString
	case AddCheckedExpr:
		return AddCheckedExprString
	case AndExpr:
		return AndExprString
	case AndAlsoExpr:
		return AndAlsoExprString
	case ConstantExpr:
		return ConstantExprString
	case DivideExpr:
		return DivideExprString
	case EqualExpr:
		return EqualExprString
	case ExclusiveOrExpr:
		return ExclusiveOrExprString
	case GreaterThanExpr:
		return GreaterThanExprString
	case GreaterThanOrEqualExpr:
		return GreaterThanOrEqualExprString
	case LeftShiftExpr:
		return LeftShiftExprString
	case LessThanExpr:
		return LessThanExprString
	case LessThanOrEqualExpr:
		return LessThanOrEqualExprString
	case ModuloExpr:
		return ModuloExprString
	case MultiplyExpr:
		return MultiplyExprString
	case MultiplyCheckedExpr:
		return MultiplyCheckedExprString
	case NegateExpr:
		return NegateExprString
	case UnaryPlusExpr:
		return UnaryPlusExprString
	case NotExpr:
		return NotExprString
	case NotEqualExpr:
		return NotEqualExprString
	case OrExpr:
		return OrExprString
	case OrElseExpr:
		return OrElseExprString
	case ParameterExpr:
		return ParameterExprString
	case PowerExpr:
		return PowerExprString
	case RightShiftExpr:
		return RightShiftExprString
	case SubtractExpr:
		return SubtractExprString
	case SubtractCheckedExpr:
		return SubtractCheckedExprString
	default:
		return UnknownExprString
	}
}

type Expression interface {
	String() string
	Kind() reflect.Kind
	Type() ExpressionType
	NodeType() string
}
