/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	return removeElement1(nums, val)
}

// 双指针
func removeElement1(nums []int, val int) int {
	var fast = 0
	var slow = 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// @lc code=end

