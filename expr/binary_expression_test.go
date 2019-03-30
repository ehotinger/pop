package expr

import (
	"reflect"
	"testing"
)

func TestNewBinaryExpression(t *testing.T) {
	for _, test := range []struct {
		nodeType ExpressionType
		left     Expression
		right    Expression
		kind     reflect.Kind
		expected string
	}{
		{
			GreaterThanExpr,
			NewConstantExpression(10, reflect.Int32),
			NewConstantExpression(15, reflect.Int32),
			reflect.Bool,
			"(10 > 15)",
		},
	} {
		expr := NewBinaryExpression(test.nodeType, test.left, test.right, test.kind)
		if actual := expr.ToString(); actual != test.expected {
			t.Fatalf("expected %v but got %v", test.expected, expr.ToString())
		}
	}
}

func TestBinaryExpressionToString(t *testing.T) {
	for _, test := range []struct {
		expr     *BinaryExpression
		expected string
	}{
		{
			nil,
			"<nil>",
		},
		{
			NewBinaryExpression(
				UnknownExpr,
				NewConstantExpression(1, reflect.Int32),
				NewConstantExpression(2, reflect.Int32),
				reflect.Bool),
			"UnknownExpr (1, 2)",
		},
		{
			NewBinaryExpression(
				LessThanOrEqualExpr,
				NewConstantExpression(5, reflect.Int32),
				NewConstantExpression(6, reflect.Int32),
				reflect.Bool),
			"(5 <= 6)",
		},
	} {
		if actual := test.expr.ToString(); actual != test.expected {
			t.Fatalf("expected %v but got %v", test.expected, test.expr.ToString())
		}
	}
}
