package main

/*
题目：
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
*/

/*
方法一： 首尾双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：28 ms	内存消耗：6.3 MB
*/
func exchangeFunc1(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	start, end := 0, len(nums)-1
	for start != end {
		for start != end && !isEven(nums[start]) {
			start++
		}
		for start != end && isEven(nums[end]) {
			end--
		}
		if start < end {
			nums[start], nums[end] = nums[end], nums[start]
		}
	}
	return nums
}

func isEven(a int) bool {
	return a&1 == 0
}

/*
方法二： 快慢双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：36 ms	内存消耗：6.3 MB
*/
func exchangeFunc2(nums []int) []int {
	fast, slow := 0, 0
	for fast < len(nums) {
		if !isEven(nums[fast]) {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
	return nums
}
