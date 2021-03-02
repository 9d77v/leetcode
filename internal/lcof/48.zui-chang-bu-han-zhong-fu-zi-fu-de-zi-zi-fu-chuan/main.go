package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

提示：

s.length <= 40000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof
*/

/*
方法一：滑动窗口,存储数量
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：8 ms	内存消耗：3.1 MB
*/
func lengthOfLongestSubstring(s string) int {
	left, chMap, max := 0, make(map[byte]int, 0), 0
	for right := range s {
		chMap[s[right]]++
		for chMap[s[right]] > 1 {
			chMap[s[left]]--
			left++
		}
		max = Max(max, right-left+1)
	}
	return max
}

/*
方法二：滑动窗口,存储位置
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：8 ms	内存消耗：3.1 MB
*/
func lengthOfLongestSubstringFunc2(s string) int {
	left, chMap, max := 0, make(map[byte]int, 0), 0
	for right := range s {
		if pos, ok := chMap[s[right]]; ok {
			left = Max(left, pos+1)
		}
		chMap[s[right]] = right
		max = Max(max, right-left+1)
	}
	return max
}
