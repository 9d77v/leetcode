package main

/*
题目：
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

你可以按任意顺序返回答案。

提示：

2 <= nums.length <= 10^3
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
只会存在一个有效答案

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
*/

/*
方法一：暴力
时间复杂度：О(n^2)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：3 MB
*/
func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/*
方法二：哈希表
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：3 MB
*/
func twoSumFunc2(nums []int, target int) []int {
	numMap := make(map[int]int, 0)
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return []int{}
}
