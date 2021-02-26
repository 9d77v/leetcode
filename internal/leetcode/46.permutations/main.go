package main

/*
题目：全排列
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations/
*/

/*
方法一：回溯法
时间复杂度：О(n!)
空间复杂度：О(n^2)
运行时间：0 ms	内存消耗：2.6 MB
*/
func permute(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	var backtrack func(x int)
	backtrack = func(x int) {
		if x == n-1 {
			newNum := make([]int, n)
			copy(newNum, nums)
			result = append(result, append([]int(nil), nums...))
			return
		}
		usedMap := make(map[int]struct{}, 0)
		for i := x; i < n; i++ {
			if _, ok := usedMap[nums[i]]; !ok {
				usedMap[nums[i]] = struct{}{}
				nums[i], nums[x] = nums[x], nums[i]
				backtrack(x + 1)
				nums[i], nums[x] = nums[x], nums[i]
			}
		}
	}
	backtrack(0)
	return result
}
