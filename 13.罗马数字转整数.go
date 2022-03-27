/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	return romanToInt1(s)
}

var symbolValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// 模拟
func romanToInt1(s string) int {
	var result int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && symbolValues[s[i]] < symbolValues[s[i+1]] {
			result -= symbolValues[s[i]]
		} else {
			result += symbolValues[s[i]]
		}
	}
	return result
}

// @lc code=end

