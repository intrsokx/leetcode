package main

//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
//
// 返回同样按升序排列的结果链表。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//
//
// 示例 2：
//
//
//输入：head = [1,1,1,2,3]
//输出：[2,3]
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序排列
//
// Related Topics 链表
// 👍 643 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	pre := dummy

	delV := 0
	for pre != nil && pre.Next != nil {
		//需要记录删除的值
		//head = cur.Next
		node := pre.Next
		if node.Next != nil && node.Val == node.Next.Val {
			delV = node.Val
			//node需要删除,并且删除node后面与node值相同的节点
			for node != nil && node.Val == delV {
				pre.Next = node.Next
				node = node.Next
			}
		} else {
			pre = pre.Next
		}
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
