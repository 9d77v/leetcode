package main

/*
题目：众数 II
给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。


提示：

1 <= nums.length <= 5 * 10^4
-10^9 <= nums[i] <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element-ii
*/

/*
方法一: 摩尔投票法
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：12 ms	内存消耗：4.9 MB
*/
func majorityElement(nums []int) []int {
	vote1, vote2, x1, x2 := 0, 0, nums[0], nums[0]
	for _, num := range nums {
		if num == x1 {
			vote1++
		} else if num == x2 {
			vote2++
		} else if vote1 == 0 {
			x1 = num
			vote1++
		} else if vote2 == 0 {
			x2 = num
			vote2++
		} else {
			vote1--
			vote2--
		}
	}
	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if num == x1 {
			cnt1++
		}
		if num == x2 {
			cnt2++
		}
	}
	result := make([]int, 0)
	if cnt1 > len(nums)/3 {
		result = append(result, x1)
	}
	if cnt2 > len(nums)/3 && x1 != x2 {
		result = append(result, x2)
	}
	return result
}
