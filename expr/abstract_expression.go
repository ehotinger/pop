package expr

import "reflect"

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

func (e *AbstractExpression) Type() ExpressionType {
	return e.nodeType
}
