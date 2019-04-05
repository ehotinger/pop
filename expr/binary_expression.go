package expr

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	errLeftNil  = errors.New("left cannot be nil")
	errRightNil = errors.New("right cannot be nil")
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
			kind:     kind,
		},
	}
}

func (e *BinaryExpression) Left() Expression {
	return e.left
}

func (e *BinaryExpression) Right() Expression {
	return e.right
}

func (e *BinaryExpression) Kind() reflect.Kind {
	return e.self.kind
}

func (e *BinaryExpression) Type() ExpressionType {
	return e.self.nodeType
}

func (e *BinaryExpression) NodeType() string {
	return "BinaryExpression"
}

func (e *BinaryExpression) String() string {
	if e == nil {
		return "<nil>"
	}
	operator := e.GetOperator()
	if operator == "" {
		return fmt.Sprintf("%v (%v, %v)", e.self, e.left, e.right)
	}
	return fmt.Sprintf("(%v %v %v)", e.left, operator, e.right)
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
		if e.self.Kind() != reflect.Bool {
			return "&"
		}
		return andIdentifier
	case AndAlsoExpr:
		return "&&"
	case OrExpr:
		if e.self.Kind() != reflect.Bool {
			return "|"
		}
		return orIdentifier
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

func CreateEqual(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	return NewBinaryExpression(EqualExpr, left, right, reflect.Bool), nil
}

func CreateNotEqual(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	return NewBinaryExpression(NotEqualExpr, left, right, reflect.Bool), nil
}

func CreateGreaterThan(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	return NewBinaryExpression(GreaterThanExpr, left, right, reflect.Bool), nil
}

func CreateGreaterThanOrEqual(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	return NewBinaryExpression(GreaterThanOrEqualExpr, left, right, reflect.Bool), nil
}

func CreateLessThan(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	return NewBinaryExpression(LessThanExpr, left, right, reflect.Bool), nil
}

func CreateLessThanOrEqual(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	return NewBinaryExpression(LessThanOrEqualExpr, left, right, reflect.Bool), nil
}

func CreateAdd(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	// TODO: String concat
	if IsArithmetic(left.Kind()) && IsArithmetic(right.Kind()) {
		return NewBinaryExpression(AddExpr, left, right, left.Kind()), nil
	}
	return nil, fmt.Errorf("invalid expression, left or right isn't arithmetic: %v, %v", left.Kind(), right.Kind())
}

func CreateSubtract(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	return NewBinaryExpression(SubtractExpr, left, right, left.Kind()), nil
}

func CreateMultiply(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	return NewBinaryExpression(MultiplyExpr, left, right, left.Kind()), nil
}

func CreateDivide(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	return NewBinaryExpression(DivideExpr, left, right, left.Kind()), nil
}

func CreateModulus(left Expression, right Expression) (Expression, error) {
	if err := validateLeftAndRight(left, right); err != nil {
		return nil, err
	}
	return NewBinaryExpression(ModuloExpr, left, right, left.Kind()), nil
}

func validateLeftAndRight(left Expression, right Expression) error {
	if left == nil {
		return errLeftNil
	}
	if right == nil {
		return errRightNil
	}
	return nil
}
