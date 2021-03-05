package main

/*
题目：和为s的两个数字
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可

来源：力扣（LeetCode）
https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
*/

/*
方法一：哈希表
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：328 ms	内存消耗：11.6 MB
*/
func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{nums[j], num}
		}
		numMap[num] = i
	}
	return []int{}
}

/*
方法二：双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：280 ms	内存消耗：8.8 MB
*/
func twoSumFunc2(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		curSum := nums[left] + nums[right]
		if curSum == target {
			return []int{nums[left], nums[right]}
		}
		if curSum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
