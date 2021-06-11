package treeModel

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func GenerateNode(val interface{}, left, right *TreeNode) *TreeNode {
	node := &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
	return node
}
