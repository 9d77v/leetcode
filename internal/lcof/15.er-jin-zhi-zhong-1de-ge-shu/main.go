package main

/*
题目：二进制中1的个数
请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。

提示：
输入必须是长度为 32 的 二进制串 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof
*/

/*
方法一：逐位判断
时间复杂度：О(logn)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：1.9 MB
*/
func hammingWeight(num uint32) int {
	var res uint32 = 0
	for num != 0 {
		res += num & 1
		num >>= 1
	}
	return int(res)
}

/*
方法一：巧用n&(n-1)
时间复杂度：О(m)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：1.9 MB
*/
func hammingWeightFunc2(num uint32) int {
	res := 0
	for num != 0 {
		res++
		num &= num - 1
	}
	return res
}
