package main

import "fmt"

//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链
//表节点，返回 反转后的链表 。
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//
//
// 示例 2：
//
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
//
//
//
//
// 提示：
//
//
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
//
//
// 进阶： 你可以使用一趟扫描完成反转吗？
// Related Topics 链表
// 👍 930 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//反转从位置m到位置n的链表
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//先遍历到m处
	dummy := &ListNode{Next: head}
	cur := dummy
	var pre *ListNode
	i := 0

	for i < m {
		pre = cur
		cur = cur.Next
		i++
	}
	//cur 目前在m的位置
	//connect, tail 用于中间节点连接
	connect := pre
	tail := cur

	for j := i; j <= n; j++ {
		temp := cur.Next

		cur.Next = pre
		pre = cur
		cur = temp
	}
	tail.Next = cur
	if connect == nil {
		dummy.Next = pre
	} else {
		connect.Next = pre
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	//fmt.Println(reverseBetween(list, 1, 5))
	ret := reverseBetween(list, 2, 4)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
