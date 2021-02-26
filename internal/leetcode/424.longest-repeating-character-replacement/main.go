package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：替换后的最长重复字符
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。

注意：字符串长度 和 k 不会超过 104。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-repeating-character-replacement
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：О(∣Σ∣)
运行时间：0 ms	内存消耗：2.4 MB
*/
func characterReplacement(s string, k int) int {
	left, maxCnt, chArr := 0, 0, [26]int{}
	for right, ch := range s {
		chArr[ch-'A']++
		maxCnt = Max(maxCnt, chArr[ch-'A'])
		if right-left+1 > maxCnt+k {
			chArr[s[left]-'A']--
			left++

		}
	}
	return len(s) - left
}
