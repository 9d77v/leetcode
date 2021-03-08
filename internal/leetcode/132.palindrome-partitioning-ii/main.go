package main

import (
	"math"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：分割回文串 II

给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。

返回符合要求的 最少分割次数 。

提示：

1 <= s.length <= 2000
s 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning-ii/
*/

/*
方法一：动态规划
时间复杂度：О(n*2)
空间复杂度：О(n^2)
运行时间：8 ms	内存消耗：5 MB
*/
func minCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	max := 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			if dp[i][j] && j-i+1 > max {
				max = j - i + 1
			}
		}
	}
	if max == n {
		return 0
	}
	f := make([]int, n)
	for i := 0; i < n; i++ {
		if dp[0][i] {
			continue
		}
		f[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if dp[j+1][i] {
				f[i] = Min(f[i], f[j]+1)
			}
		}
	}
	return f[n-1]
}
