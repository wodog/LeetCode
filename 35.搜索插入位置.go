/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	return searchInsert1(nums, target)
}

// 第一次解
func searchInsert1(nums []int, target int) int {
	if len(nums) == 0 { // 边界case
		return 0
	}

	// 如果不存在
	if target < nums[0] { // 比第一个小
		return 0
	} else if target > nums[len(nums)-1] { // 比最后一个大
		return len(nums)
	} else { // 介于两个数之间
		for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
			if nums[i] > target && nums[i-1] < target {
				return i
			}
		}
	}

	return -1
}

// @lc code=end

