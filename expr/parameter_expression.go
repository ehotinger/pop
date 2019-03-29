package expr

import "reflect"

type ParameterExpression struct {
	self *AbstractExpression
	name string
}

func NewParameterExpression(name string, kind reflect.Kind) *ParameterExpression {
	return &ParameterExpression{
		self: &AbstractExpression{
			nodeType: ParameterExpr,
			Kind:     kind,
		},
		name: name,
	}
}

func (e *ParameterExpression) Name() string {
	return e.name
}

func (e *ParameterExpression) ToString() string {
	if e.name == "" {
		return "<param>"
	}
	return e.name
}
