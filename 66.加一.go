/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			continue
		} else {
			digits[i] += 1
			break
		}
	}

	if digits[0] == 0 {
		arr := make([]int, len(digits)+1)
		arr[0] = 1
		return arr
	}

	return digits
}

// @lc code=end

