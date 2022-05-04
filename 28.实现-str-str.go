/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	return strStr1(haystack, needle)
}

// 字符串匹配，后者完全匹配前者
func strStr1(haystack string, needle string) int {
	if needle == "" { // 匹配的串是空时返回0
		return 0
	}
	if len(needle) <= len(haystack) { // 匹配的串长度要求小于等于原始串
		for i := 0; i < len(haystack); i++ {
			if i+len(needle) > len(haystack) { // 判断模式串的越界情况
				break
			}
			var matchstr = haystack[i : i+len(needle)] // 模式串
			if matchstr == needle {                    // 比较模式串与匹配串是否相等
				return i
			}
		}
	}
	return -1
}

// @lc code=end

