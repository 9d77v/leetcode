package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。
提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring/
*/

/*
方法一：暴力解法
时间复杂度：O(n^3)
空间复杂度：O(1)
运行时间：144 ms	内存消耗：2.7 MB
*/
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	left, max := 0, 1
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if j-i+1 > max && IsPalidrome(s, i, j) {
				max = j - i + 1
				left = i
			}
		}
	}
	return s[left : left+max]
}

/*
方法二：中心扩散
时间复杂度：O(n^2)
空间复杂度：O(1)
运行时间：4 ms	内存消耗：2.6 MB
*/
func longestPalindromeFunc2(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	left, max := 0, 1
	for i := 0; i < n-1; i++ {
		oddLen, evenLen := expandCenter(s, i, i), expandCenter(s, i, i+1)
		curMax := Max(oddLen, evenLen)
		if curMax > max {
			max = curMax
			left = i - (max-1)>>1
		}
	}
	return s[left : left+max]
}

func expandCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1
}

/*
方法三：动态规划
时间复杂度：O(n^2)
空间复杂度：O(n^2)
运行时间：176 ms	内存消耗：7 MB
*/
func longestPalindromeFunc3(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	left, max := 0, 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			if dp[i][j] && j-i+1 > max {
				max = j - i + 1
				left = i
			}
		}
	}
	return s[left : left+max]
}

/*
方法四：Manacher
时间复杂度：O(n)
空间复杂度：O(n^2)
运行时间：32 ms	内存消耗：7.3 MB
*/
func longestPalindromeFunc4(s string) string {
	start, end := 0, -1
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	t += "#"
	s = t
	armLen := []int{}
	right, j := -1, -1
	for i := 0; i < len(s); i++ {
		var curArmLen int
		if right >= i {
			iSym := j*2 - i
			minArmLen := Min(armLen[iSym], right-i)
			curArmLen = expand(s, i-minArmLen, i+minArmLen)
		} else {
			curArmLen = expand(s, i, i)
		}
		armLen = append(armLen, curArmLen)
		if i+curArmLen > right {
			j = i
			right = i + curArmLen
		}
		if curArmLen*2+1 > end-start {
			start = i - curArmLen
			end = i + curArmLen
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			ans += string(s[i])
		}
	}
	return ans
}

func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) / 2
}
