package main

/*
题目： 分割回文串

给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning/
*/

/*
方法一：回溯+动态规划预处理
时间复杂度：О(n*2^n)
空间复杂度：О(n^2)
运行时间：340 ms	内存消耗：23.1 MB
*/
func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	result := [][]string{}
	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			result = append(result, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return result
}
