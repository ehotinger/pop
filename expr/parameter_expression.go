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
			kind:     kind,
		},
		name: name,
	}
}

func (e *ParameterExpression) Name() string {
	return e.name
}

func (e *ParameterExpression) Kind() reflect.Kind {
	return e.self.kind
}

func (e *ParameterExpression) Type() ExpressionType {
	return e.self.nodeType
}

func (e *ParameterExpression) ToString() string {
	if e.name == "" {
		return "<param>"
	}
	return e.name
}
