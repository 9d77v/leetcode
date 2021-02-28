package main

import "strconv"

/*
题目：
数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。

限制：

0 <= n < 2^31

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof
*/

/*
方法一: 找数学规律
时间复杂度：О(logN)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：1.9 MB
*/
func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit++
		count = 9 * start * digit
	}
	num := start + (n-1)/digit
	result, _ := strconv.Atoi(string(strconv.Itoa(num)[(n-1)%digit]))
	return result
}
