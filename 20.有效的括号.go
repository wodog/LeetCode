/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	return isValid1(s)
}

// 栈
func isValid1(s string) bool {
	var pairs map[byte]byte = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []byte
	for i := 0; i < len(s); i++ {
		if val, exist := pairs[s[i]]; exist {
			if len(stack) != 0 && stack[len(stack)-1] == val {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

// @lc code=end

