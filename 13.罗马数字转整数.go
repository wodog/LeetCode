/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	return romanToInt1(s)
}

func romanToInt1(s string) int {
	var m map[byte]int = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && m[s[i+1]] > m[s[i]] {
			result -= m[s[i]]
		} else {
			result += m[s[i]]
		}
	}
	return result
}

// @lc code=end

