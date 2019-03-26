package expr

import "reflect"

type ParameterExpression struct {
	self *AbstractExpression
	name string
}

func NewParameterExpression(name string, staticType reflect.Type) *ParameterExpression {
	return &ParameterExpression{
		self: &AbstractExpression{
			nodeType: ParameterExpr,
			Type:     staticType,
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
