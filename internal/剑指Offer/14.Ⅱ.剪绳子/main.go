package main

/*
题目：
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1] 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

提示：
2 <= n <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof
*/

/*
方法一：数学方法
时间复杂度：О(1)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：1.9 MB

当绳子尽可能切成3时乘积最大
*/
func cuttingRopeFunc1(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	t := n % 3
	m := n / 3
	if t == 0 {
		t++
	} else if t == 1 {
		m--
		t += 3
	}
	res := t
	for i := 0; i < m; i++ {
		res = 3 * res
		if res > 1000000007 {
			res %= 1000000007
		}
	}
	return res
}
