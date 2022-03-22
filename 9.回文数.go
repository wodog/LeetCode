/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	return isPalindrome(x)
}

// 整数转为字符串
func isPalindrome1(x int) bool {
	s := strconv.Itoa(x)

	// 双指针
	left := 0
	right := len(s) - 1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end

