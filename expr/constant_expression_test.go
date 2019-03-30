package expr

import (
	"reflect"
	"testing"
)

func TestToString(t *testing.T) {
	for _, test := range []struct {
		expr     *ConstantExpression
		expected string
	}{
		{
			nil,
			"<nil>",
		},
		{
			&ConstantExpression{},
			"<nil>",
		},
		{
			&ConstantExpression{
				self: &AbstractExpression{
					nodeType: ConstantExpr,
					kind:     reflect.String,
				},
				value: "HelloWorld!",
			},
			`"HelloWorld!"`,
		},
		{
			&ConstantExpression{
				self: &AbstractExpression{
					nodeType: ConstantExpr,
					kind:     reflect.Int32,
				},
				value: int32(300),
			},
			"value(300)",
		},
		{
			&ConstantExpression{
				self: &AbstractExpression{
					nodeType: ConstantExpr,
					kind:     reflect.Int32,
				},
				value: int64(300),
			},
			"300",
		},
	} {
		if actual := test.expr.ToString(); actual != test.expected {
			t.Fatalf("expected %v but got %v", test.expected, actual)
		}
	}
}
