package main

import "fmt"

//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，po
//s 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
//
// 说明：不允许修改给定的链表。
//
// 进阶：
//
//
// 你是否可以使用 O(1) 空间解决此题？
//
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [3,2,0,-4], pos = 1
//输出：返回索引为 1 的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2], pos = 0
//输出：返回索引为 0 的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
// 示例 3：
//
//
//
//
//输入：head = [1], pos = -1
//输出：返回 null
//解释：链表中没有环。
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围在范围 [0, 104] 内
// -105 <= Node.val <= 105
// pos 的值为 -1 或者链表中的一个有效索引
//
// Related Topics 哈希表 链表 双指针
// 👍 1045 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//求环的入口节点:设从链头到入口处的距离为x, 从入口到相遇处为y，从相遇处到入口为z
//2*(x+y) = x+y + n(y+z)
//x+y = (n-1)*(y+z) + y + z
//x = (n-1)*(y+z) + z

/**
tips:
指针比较时直接比较对象，不要用值比较，链表中有可能存在重复值情况
第一次相交后，快指针从头结点开始，慢指针则从下一个节点开始
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head.Next, head

	for fast != nil && fast.Next != nil {
		if fast == slow {
			fast = head
			slow = slow.Next
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var meet *ListNode
	tail := &ListNode{
		Val:  4,
		Next: nil,
	}
	meet = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  0,
			Next: tail,
		},
	}
	tail.Next = meet

	list := &ListNode{
		Val:  3,
		Next: meet,
	}

	ret := detectCycle(list)
	fmt.Println(ret == meet)
}
