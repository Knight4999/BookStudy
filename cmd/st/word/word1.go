package word

// 包word提供了文字游戏相关函数工具

//IsPalindrom 判断一个字符串是否返回文字字符串
func IsPalindrom(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
