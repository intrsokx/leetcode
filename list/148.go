package main

import "fmt"

//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
// 进阶：
//
//
// 你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
//
//
//
// 示例 1：
//
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
//
// 示例 2：
//
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 5 * 104] 内
// -105 <= Node.val <= 105
//
// Related Topics 链表 双指针 分治 排序 归并排序
// 👍 1187 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
	链表的排序，因为不能像数组那样交换，所以像一些以交换为主的方法是不适应于链表的。
比如有快速排序算法、冒泡排序，插入排序等
*/
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

//思路：归并排序，找中点再合并（当一个链表只有一个节点时，它必然是有序的）
func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//split
	mid := findMidNode(head)

	tail := mid.Next
	mid.Next = nil

	left := mergeSort(head)
	right := mergeSort(tail)

	ret := merge2list(left, right)
	return ret
}

func findMidNode(head *ListNode) *ListNode {
	//check (虽然在mergeSort中做了判断，但是为了防止上层调用者忘记check导致在当前函数中panic，所以这块再check一次)
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

func merge2list(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			tail = l1

			l1 = l1.Next
		} else {
			tail.Next = l2
			tail = l2

			l2 = l2.Next
		}
	}

	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
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
			Val: 4,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	fmt.Println(sortList(list))
}
