package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：О(∣Σ∣)
运行时间：16 ms	内存消耗：2.8 MB
*/
func lengthOfLongestSubstring(s string) int {
	left, maxCnt, chMap := 0, 0, make(map[byte]int, 0)
	for right, ch := range s {
		if _, ok := chMap[byte(ch)]; ok {
			newLeft := chMap[byte(ch)] + 1
			for i := left; i < newLeft; i++ {
				delete(chMap, s[i])
			}
			left = newLeft
		}
		chMap[byte(ch)] = right
		maxCnt = Max(maxCnt, right-left+1)
	}
	return maxCnt
}
