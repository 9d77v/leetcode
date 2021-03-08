package main

/*
题目：
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break
*/

/*
方法一：动态规划
时间复杂度：О(n^2)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2.1 MB
*/
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordMap[v] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
