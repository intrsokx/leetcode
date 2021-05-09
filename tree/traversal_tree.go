package main

import "fmt"

//遍历二叉树(前序、中序、后序)

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
	preOrderTraversalV1(root.Left)
	fmt.Printf("%v\t", root.Val)
	preOrderTraversalV1(root.Right)
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
	preOrderTraversalV1(root.Left)
	preOrderTraversalV1(root.Right)
	fmt.Printf("%v\t", root.Val)

}
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
		if tmp.Right == nil || tmp.Right == lastVisited {
			stack = stack[:len(stack)-1]
			lastVisited = tmp
			ret = append(ret, tmp.Val)
		} else {
			dummy = tmp.Right
		}
	}

	return ret
}
