// 包word提供了文字游戏相关的工具函数
package word

// IsPallindrome判断一个字符串是否是回文字符串
// (Our first attempt.)
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
