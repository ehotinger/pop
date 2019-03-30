package expr

type BinaryVisitor struct {
	node *BinaryExpression
}

func NewBinaryVisitor(node *BinaryExpression) *BinaryVisitor {
	return &BinaryVisitor{
		node: node,
	}
}
