package main

import (
	"sort"
	"strconv"
)

/*
题目：把数组排成最小的数
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

提示:

0 < nums.length <= 100
说明:

输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof
*/

/*
方法一: 排序x+y>y+x
时间复杂度：О(logN)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.6 MB
*/
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, _ := strconv.Atoi(strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]))
		b, _ := strconv.Atoi(strconv.Itoa(nums[j]) + strconv.Itoa(nums[i]))
		return a < b
	})
	numBytes := make([]byte, 0)
	for _, num := range nums {
		numBytes = append(numBytes, []byte(strconv.Itoa(num))...)
	}
	return string(numBytes)
}
