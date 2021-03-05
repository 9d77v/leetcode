package main

/*
题目：
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

限制：

2 <= nums.length <= 10000

来源：力扣（LeetCode）
https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
*/

/*
方法一：分组异或
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：28 ms	内存消耗：6.3 MB
*/
func singleNumbers(nums []int) []int {
	k := 0
	for _, num := range nums {
		k ^= num
	}
	mask := 1
	for mask&k == 0 {
		mask <<= 1
	}
	a, b := 0, 0
	for _, num := range nums {
		if mask&num != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
