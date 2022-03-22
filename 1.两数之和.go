/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	return twoSum2(nums, target)
}

// 暴力枚举
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 哈希表
func twoSum2(nums []int, target int) []int {
	var hashmap map[int]int = make(map[int]int)
	for i := range nums {
		if val, ok := hashmap[target-nums[i]]; ok {
			return []int{val, i}
		}
		hashmap[nums[i]] = i
	}
	return nil
}

// @lc code=end
