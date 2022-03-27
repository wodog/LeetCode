/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return mergeTwoLists3(list1, list2)
}

// 递归
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list2.Next, list1)
		return list2
	}
}

// 迭代
func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	var prehead = &ListNode{}
	var prev *ListNode = prehead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}

	if list1 == nil {
		prev.Next = list2
	}
	if list2 == nil {
		prev.Next = list1
	}

	return prehead.Next
}

// 合并成数组再排序，转成链表
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	var list *ListNode = &ListNode{}
	var head = list
	var arr []int
	for list1 != nil {
		arr = append(arr, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil {
		arr = append(arr, list2.Val)
		list2 = list2.Next
	}

	popSort(arr)

	for _, val := range arr {
		list.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		list = list.Next
	}

	return head.Next
}

func popSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
	}
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// @lc code=end

