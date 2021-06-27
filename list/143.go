package main

//给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
//将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 示例 1:
//
// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
//
// 示例 2:
//
// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
// Related Topics 栈 递归 链表 双指针
// 👍 604 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//思路：找到中点断开，然后翻转后半部分，再挨个合并
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mid := findMidNodeV2(head)

	tail := mid.Next
	mid.Next = nil

	tail = reverseListV3(tail)

	ret := conncet2list(head, tail)

	head = ret
}

func findMidNodeV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head.Next, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverseListV3(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre

		pre = cur
		cur = tmp
	}

	return pre
}

func conncet2list(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	//len(l1) == len(l2) || len(l1) == len(l2)+1

	for l2 != nil {
		tmp1 := l1.Next
		tail.Next = l1
		tail = l1

		tmp2 := l2.Next
		tail.Next = l2
		tail = l2

		l1 = tmp1
		l2 = tmp2
	}

	if l1 != nil {
		tail.Next = l1
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

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

	reorderList(list)
}
