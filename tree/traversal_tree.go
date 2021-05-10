package main

import "fmt"

//遍历二叉树(前序、中序、后序)
//type TreeNode struct {
//	Val   interface{}
//	Left  *TreeNode
//	Right *TreeNode
//}


//recursion
func preOrderTraversalV1(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%v\t", root.Val)
	preOrderTraversalV1(root.Left)
	preOrderTraversalV1(root.Right)
}
func preOrderTraversalV2(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	ret := make([]interface{}, 0)
	dummy := root
	stack := make([]*TreeNode, 0)

	for dummy != nil || len(stack) > 0 {
		for dummy != nil {
			ret = append(ret, dummy.Val)
			stack = append(stack, dummy.Right)
			dummy = dummy.Left
		}
		//程序走到这块，stack中一定有数据，然后pop
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		dummy = n
	}
	return ret
}

func inOrderTraversalV1(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversalV1(root.Left)
	fmt.Printf("%v\t", root.Val)
	inOrderTraversalV1(root.Right)
}
func inOrderTraversalV2(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	ret := make([]interface{}, 0)
	dummy := root
	stack := make([]*TreeNode, 0)

	for dummy != nil || len(stack) > 0 {
		for dummy != nil {
			stack = append(stack, dummy)
			dummy = dummy.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, n.Val)
		dummy = n.Right
	}
	return ret
}

func postOrderTraversalV1(root *TreeNode) {
	if root == nil {
		return
	}
	postOrderTraversalV1(root.Left)
	postOrderTraversalV1(root.Right)
	fmt.Printf("%v\n", root.Val)

}
//左->右->根
//当左右都访问后才能访问根
func postOrderTraversalV2(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	ret := make([]interface{}, 0)
	dummy := root
	stack := make([]*TreeNode, 0)
	var lastVisited *TreeNode

	for dummy != nil || len(stack) > 0 {
		for dummy != nil {
			stack = append(stack, dummy)
			dummy = dummy.Left
		}

		tmp := stack[len(stack)-1]
		//到这一步tmp.Left一定为nil，那么就需要判断temp的右节点是否也为nil，如果为nil，这时候就可以访问tmp节点了
		//如果tmp的右节点不为nil，则需要判断这个右节点是否访问过，如果没有访问过，
		//则递归的处理该节点（就想重新处理一个root节点那样:dummy=tmp.Right）；
		//如果访问过了，则当前tmp节点出栈，访问该节点，并设置访问标志（lastVisited）
		if tmp.Right == nil || tmp.Right == lastVisited {
			//
			stack = stack[:len(stack)-1]
			lastVisited = tmp
			ret = append(ret, tmp.Val)
		} else {
			dummy = tmp.Right
		}
	}

	return ret
}

//【1， 2， 3， 4， 5， nil, nil, nil, nil, nil, nil】
func GenerateTree(values []interface{}) *TreeNode {
	var dummy *TreeNode
	if len(values) == 0 {
		return dummy
	}
	
	for _, val := range values {
		_ = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	return dummy
}

//go run  tree.go traversal_tree.go
func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(preOrderTraversalV2(root))
	fmt.Println(inOrderTraversalV2(root))
	//postOrderTraversalV1(root)
	fmt.Println(postOrderTraversalV2(root))
}