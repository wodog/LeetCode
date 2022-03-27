/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	return deleteDuplicates1(head)
}

// 双指针
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var slow = head
	var fast = head

	for fast != nil {
		if fast.Val != slow.Val {
			slow = slow.Next
			slow.Val = fast.Val
		}
		fast = fast.Next
	}

	slow.Next = nil
	return head
}

// @lc code=end

