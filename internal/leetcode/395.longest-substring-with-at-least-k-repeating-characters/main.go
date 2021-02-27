package main

import (
	"strings"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：至少有K个重复字符的最长子串
找到给定字符串（由小写字符组成）中的最长子串 T ， 要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string/
*/

/*
方法一：分治
时间复杂度： О(n)
空间复杂度：О(|Σ|)
运行时间：0 ms	内存消耗：2.3 MB
*/
func longestSubstring(s string, k int) (ans int) {
	if s == "" {
		return
	}
	chArr := [26]int{}
	for _, ch := range s {
		chArr[ch-'a']++
	}
	var split byte
	for i, c := range chArr {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}
	for _, subStr := range strings.Split(s, string(split)) {
		ans = Max(ans, longestSubstring(subStr, k))
	}
	return
}
