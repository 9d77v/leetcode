package main

import "math/bits"

/*
题目：
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

进阶:

给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
要求算法的空间复杂度为O(n)。
你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/counting-bits
*/

/*
方法一：直接计算，使用内置函数
时间复杂度：О(k*num)
空间复杂度：O(1)
运行时间：4 ms	内存消耗：4.6 MB
*/
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		result[i] = bits.OnesCount(uint(i))
	}
	return result
}

/*
方法二：直接计算，不使用内置函数
时间复杂度：О(k*num)
空间复杂度：O(1)
运行时间：4 ms	内存消耗：4.6 MB
*/
func countBitsFunc2(num int) []int {
	result := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		result[i] = onesCount(i)
	}
	return result
}

func onesCount(num int) (count int) {
	for ; num > 0; num &= num - 1 {
		count++
	}
	return
}

/*
方法三：动态规划，最高有效位
时间复杂度：О(num)
空间复杂度：O(1)
运行时间：4 ms	内存消耗：4.6 MB
*/
func countBitsFunc3(num int) []int {
	result := make([]int, num+1)
	hightBit := 0
	for i := 1; i < num+1; i++ {
		if i&(i-1) == 0 {
			hightBit = i
		}
		result[i] = result[i-hightBit] + 1
	}
	return result
}

/*
方法四：动态规划，最低有效位
时间复杂度：О(num)
空间复杂度：O(1)
运行时间：4 ms	内存消耗：4.6 MB
*/
func countBitsFunc4(num int) []int {
	result := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		result[i] = result[i>>1] + i&1
	}
	return result
}

/*
方法五：动态规划，最低设置位
时间复杂度：О(num)
空间复杂度：O(1)
运行时间：4 ms	内存消耗：4.6 MB
*/
func countBitsFunc5(num int) []int {
	result := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}
