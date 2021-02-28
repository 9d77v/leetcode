package main

import (
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：数组中出现次数超过一半的数字
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
限制：

1 <= 数组长度 <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
*/

/*
方法一: 摩尔投票法
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：16 ms	内存消耗：6.1 MB
*/
func majorityElement(nums []int) int {
	vote, x := 0, nums[0]
	for _, num := range nums {
		if vote == 0 {
			x = num
		}
		if num == x {
			vote++
		} else {
			vote--
		}
	}
	return x
}

/*
方法二: 哈希表
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：24 ms	内存消耗：6.1 MB
*/
func majorityElementFunc2(nums []int) int {
	numMap := map[int]int{}
	for _, num := range nums {
		numMap[num]++
	}
	for k, v := range numMap {
		if v > len(nums)>>1 {
			return k
		}
	}
	return -1
}

/*
方法三: 排序
时间复杂度：О(nlogn)
空间复杂度：О(logn)
运行时间：20 ms	内存消耗：6 MB
*/
func majorityElementFunc3(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)>>1]
}

/*
方法四: 快排变形
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：1004 ms	内存消耗：8.8 MB
*/
func majorityElementFunc4(nums []int) int {
	n := len(nums)
	k := n >> 1
	QuickSelect(nums, k, 0, n-1)
	return nums[k]
}
