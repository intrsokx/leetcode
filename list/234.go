package main

//请判断一个链表是否为回文链表。
//
// 示例 1:
//
// 输入: 1->2
//输出: false
//
// 示例 2:
//
// 输入: 1->2->2->1
//输出: true
//
//
// 进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
// Related Topics 栈 递归 链表 双指针
// 👍 1028 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//思路：找到中点，将链表断开成两部分，翻转后半部分，然后比较
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid := findmidnodeV3(head)

	tail := mid.Next
	mid.Next = nil

	tail = reverseListV4(tail)

	return compare2list(head, tail)
}

//返回中间节点的前一个节点
func findmidnodeV3(head *ListNode) *ListNode {
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

func reverseListV4(head *ListNode) *ListNode {
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

func compare2list(l1, l2 *ListNode) bool {
	head, tail := l1, l2

	for tail != nil && head != nil {
		if tail.Val != head.Val {
			return false
		}
		tail = tail.Next
		head = head.Next
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
