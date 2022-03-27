/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	return longestCommonPrefix1(strs)
}

// 纵向扫描
func longestCommonPrefix1(strs []string) string {
	var prefix string

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return prefix
			}
		}
		prefix += string(strs[0][i])
	}
	return prefix
}

// @lc code=end

