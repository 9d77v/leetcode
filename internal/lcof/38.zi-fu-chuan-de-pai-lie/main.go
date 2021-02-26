package main

/*
题目：字符串的排列
输入一个字符串，打印出该字符串中字符的所有排列。

你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

限制：

1 <= s 的长度 <= 8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
*/

/*
方法一：回溯法
时间复杂度：О(n!)
空间复杂度：О(n^2)
运行时间：44 ms	内存消耗：7.6 MB
*/
func permutation(s string) []string {
	n := len(s)
	result := make([]string, 0)
	c := []byte(s)
	var dfs func(x int)
	dfs = func(x int) {
		if x == n-1 {
			result = append(result, string(c))
			return
		}
		byteMap := make(map[byte]struct{}, 8)
		for i := x; i < n; i++ {
			if _, ok := byteMap[c[i]]; !ok {
				byteMap[c[i]] = struct{}{}
				c[i], c[x] = c[x], c[i]
				dfs(x + 1)
				c[i], c[x] = c[x], c[i]
			}
		}
	}
	dfs(0)
	return result
}
