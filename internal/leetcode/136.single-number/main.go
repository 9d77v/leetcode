package main

/*
题目：
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
*/

/*
方法一：异或
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：20 ms	内存消耗：6.1 MB
*/
func singleNumber(nums []int) (result int) {
	for _, num := range nums {
		result ^= num
	}
	return
}
