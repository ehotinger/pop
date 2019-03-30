package expr

import (
	"fmt"
	"reflect"
)

type ConstantExpression struct {
	self  *AbstractExpression
	value interface{}
}

func NewConstantExpression(value interface{}, kind reflect.Kind) *ConstantExpression {
	return &ConstantExpression{
		self: &AbstractExpression{
			nodeType: ConstantExpr,
			kind:     kind,
		},
		value: value,
	}
}

func (e *ConstantExpression) Kind() reflect.Kind {
	return e.self.kind
}

func (e *ConstantExpression) ToString() string {
	if e == nil || e.value == nil {
		return "<nil>"
	}
	if val, ok := e.value.(string); ok {
		return fmt.Sprintf(`"%s"`, val)
	} else if reflect.TypeOf(e.value).Kind() == e.self.Kind() { // TODO: Better reflection?
		return fmt.Sprintf("value(%v)", e.value)
	} else {
		return fmt.Sprintf("%v", e.value)
	}
}
