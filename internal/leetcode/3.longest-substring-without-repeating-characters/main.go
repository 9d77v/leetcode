package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：无重复字符的最长子串
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
		if pos, ok := chMap[byte(ch)]; ok {
			for i := left; i <= pos; i++ {
				delete(chMap, s[i])
			}
			left = pos + 1
		}
		chMap[byte(ch)] = right
		maxCnt = Max(maxCnt, right-left+1)
	}
	return maxCnt
}

/*
方法二：双指针
时间复杂度：О(n)
空间复杂度：О(∣Σ∣)
运行时间：12 ms	内存消耗：2.8 MB
*/
func lengthOfLongestSubstringFunc2(s string) int {
	right, maxCnt, chMap, n := 0, 0, make(map[byte]int, 0), len(s)
	for left := range s {
		if left != 0 {
			delete(chMap, s[left-1])
		}
		for right < n && chMap[s[right]] == 0 {
			chMap[s[right]]++
			right++
		}
		maxCnt = Max(maxCnt, right-left)
	}
	return maxCnt
}
