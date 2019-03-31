package expr

type BinaryVisitor struct {
	root *BinaryExpression
}

func NewBinaryVisitor(root *BinaryExpression) *BinaryVisitor {
	return &BinaryVisitor{
		root: root,
	}
}

func (v *BinaryVisitor) Visit() (err error) {
	var left Visitor
	left, err = CreateVisitorFromExpression(v.root.left)
	if err != nil {
		return err
	}
	err = left.Visit()
	if err != nil {
		return err
	}
	var right Visitor
	right, err = CreateVisitorFromExpression(v.root.right)
	if err != nil {
		return err
	}
	return right.Visit()
}
