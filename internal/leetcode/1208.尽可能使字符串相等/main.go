package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
给你两个长度相同的字符串，s 和 t。

将 s 中的第 i 个字符变到 t 中的第 i 个字符需要 |s[i] - t[i]| 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。

用于变更字符串的最大预算是 maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。

如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。

如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。

提示：

1 <= s.length, t.length <= 10^5
0 <= maxCost <= 10^6
s 和 t 都只含小写英文字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/get-equal-substrings-within-budget
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.8 MB
*/
func equalSubstring(s string, t string, maxCost int) int {
	left, curCost := 0, 0
	for right := range s {
		curCost += Abs(int(s[right]-'a') - int(t[right]-'a'))
		if curCost > maxCost {
			curCost -= Abs(int(s[left]-'a') - int(t[left]-'a'))
			left++
		}
	}
	return len(s) - left
}
