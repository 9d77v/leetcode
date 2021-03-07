package main

import "sort"

/*
题目：全排列II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii/
*/

/*
方法一：回溯法
时间复杂度：О(n!)
空间复杂度：О(n^2)
运行时间：0 ms	内存消耗：3.7 MB
*/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)
	path := []int{}
	visited := make([]bool, n)
	var backtrack func(int)
	backtrack = func(depth int) {
		if depth == n {
			result = append(result, append([]int(nil), path...))
			return
		}
		for i, v := range nums {
			if visited[i] || i > 0 && !visited[i-1] && v == nums[i-1] {
				continue
			}
			path = append(path, v)
			visited[i] = true
			backtrack(depth + 1)
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return result
}
