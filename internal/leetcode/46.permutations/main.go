package main

/*
题目：全排列
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations/
*/

/*
方法一：回溯法
时间复杂度：О(n*n!)
空间复杂度：О(n*n!)
运行时间：0 ms	内存消耗：2.6 MB
*/
func permute(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}
	var dfs func(int)
	dfs = func(depth int) {
		if depth == n-1 {
			result = append(result, append([]int(nil), nums...))
			return
		}
		visited := make(map[int]bool, 0)
		for i := depth; i < n; i++ {
			if !visited[nums[i]] {
				visited[nums[i]] = true
				nums[i], nums[depth] = nums[depth], nums[i]
				dfs(depth + 1)
				nums[i], nums[depth] = nums[depth], nums[i]
			}
		}
	}
	dfs(0)
	return result
}
