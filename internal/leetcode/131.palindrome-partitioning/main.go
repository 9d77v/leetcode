package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目： 分割回文串

给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning/
*/

/*
方法一：回溯
时间复杂度：О(n*2^n)
空间复杂度：О(n^2)
运行时间：320 ms	内存消耗：23.6 MB
*/
func partition(s string) [][]string {
	n := len(s)
	result := [][]string{}
	path := []string{}
	var dfs func(int)
	dfs = func(depth int) {
		if depth == n {
			result = append(result, append([]string(nil), path...))
			return
		}
		for i := depth; i < n; i++ {
			if IsPalidrome(s, depth, i) {
				path = append(path, s[depth:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return result
}

/*
方法二：回溯+记忆化
时间复杂度：О(n*2^n)
空间复杂度：О(n^2)
运行时间：328 ms	内存消耗：23.4 MB
*/
func partitionFunc2(s string) [][]string {
	n := len(s)

	memo := make([][]int8, n) //0:未计算，1：是回文，2：不是回文
	for i := range memo {
		memo[i] = make([]int8, n)
	}
	result := [][]string{}
	path := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			result = append(result, append([]string(nil), path...))
			return
		}
		for j := i; j < n; j++ {
			if memo[i][j] == 2 {
				continue
			}
			if memo[i][j] == 1 || IsPalidrome(s, i, j) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return result
}

/*
方法三：回溯+动态规划预处理
时间复杂度：О(n*2^n)
空间复杂度：О(n^2)
运行时间：340 ms	内存消耗：23.1 MB
*/
func partitionFunc3(s string) [][]string {
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
	path := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			result = append(result, append([]string(nil), path...))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return result
}
