package main

/*
题目：存在重复元素II
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-ii
*/

/*
方法一：哈希表
时间复杂度：O(n)
空间复杂度：O(n)
运行时间：16 ms	内存消耗：9.3 MB
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	dataMap := make(map[int]int, 0)
	for i, num := range nums {
		if _, ok := dataMap[num]; !ok {
			dataMap[num] = i
		} else {
			if i-dataMap[num] <= k {
				return true
			}
			dataMap[num] = i
		}
	}
	return false
}

/*
方法二：哈希表,维护k大小
时间复杂度：O(n)
空间复杂度：O(n)
运行时间：16 ms	内存消耗：6.1 MB
*/
func containsNearbyDuplicateFunc2(nums []int, k int) bool {
	dataMap := make(map[int]struct{}, 0)
	for i, num := range nums {
		if _, ok := dataMap[num]; ok {
			return true
		}
		dataMap[num] = struct{}{}
		if len(dataMap) > k {
			delete(dataMap, nums[i-k])
		}
	}
	return false
}
