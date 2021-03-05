package main

/*
题目：
在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

限制：

1 <= nums.length <= 10000
1 <= nums[i] < 2^31


来源：力扣（LeetCode）
https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
*/

/*
方法一：有限状态自动机
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：28 ms	内存消耗：6.4 MB
*/
func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, c := range nums {
		a, b = a&^c|b&c, b&^c|^a&^b&c
	}
	return ^a & b
}

/*
方法二：哈希表
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：36 ms	内存消耗：6.9 MB
*/
func singleNumberFunc2(nums []int) int {
	numMap := map[int]int{}
	for _, v := range nums {
		numMap[v]++
	}
	for i, v := range numMap {
		if v == 1 {
			return i
		}
	}
	return 0
}
