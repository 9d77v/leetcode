package main

import (
	"math"
)

/*
题目：
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof
*/

/*
方法一：二分查找
时间复杂度：О(㏒n)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：3.1 MB
*/
func minArray(numbers []int) int {
	n := len(numbers)
	if len(numbers) == 0 {
		return math.MinInt64
	}
	p, q := 0, n-1
	for p < q {
		middle := p + (q-p)/2
		if numbers[middle] > numbers[q] {
			p = middle + 1
		} else if numbers[middle] < numbers[q] {
			q = middle
		} else {
			q--
		}
	}
	return numbers[p]
}
