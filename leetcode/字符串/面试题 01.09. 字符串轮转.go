package 字符串

import "strings"

/**
题目：给定字符串s1 和 s2, 请检查s2是否可以由s1旋转而成，比如
	s1 = waterbottle
	s2 = erbottlewat

思路：只需要检查s2是不是（s1+s1）的子串即可

T= O(n) S= O(n)
*/
func isFlipedString(s1, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}
