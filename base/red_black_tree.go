package base

// 红黑树节点
type RBtreeNode struct {
	*BinaryTreeNode
	IsRed bool
}

func NewRBtreeNode(node *BinaryTreeNode) *RBtreeNode {
	return &RBtreeNode{
		BinaryTreeNode: node,
		IsRed:          true, //所有节点初始化默认为红色
	}
}

// 判断红黑树是红色
func (node *RBtreeNode) ColorIsRed() bool {
	if node == nil { //空节点代表黑色
		return false
	}
	return node.IsRed
}

// 判断红黑树是黑色
func (node *RBtreeNode) ColorIsBlack() bool {
	return !node.ColorIsRed()
}

// 给节点染色
func (node *RBtreeNode) Color(isRed bool) {
	if node == nil {
		return
	}
	node.IsRed = isRed
}

type RbTree struct {
	size int
	Root *RBtreeNode
	cmp  CompareFunc //比较
}

func (tree *RbTree) AfterRemove() {

}

func (tree *RbTree) AfterAdd() {

}
