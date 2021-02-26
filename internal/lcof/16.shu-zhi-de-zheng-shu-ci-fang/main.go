package main

/*
题目：数值的整数次方
实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

说明:

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
*/

/*
运行时间：超出时间限制
*/
func myPowFunc1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	abst := n
	if n < 0 {
		abst = -abst
	}
	res := pow(x, abst)
	if n < 0 {
		res = 1.0 / res
	}
	return res
}

func pow(x float64, n int) float64 {
	res := x
	for i := 1; i < n; i++ {
		res *= x
	}
	return res
}

/*
运行时间：0 ms	内存消耗：2.1 MB
*/
func myPowFunc2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	abst := n
	if n < 0 {
		abst = -abst
	}
	res := pow2(x, abst)
	if n < 0 {
		res = 1.0 / res
	}
	return res
}

func pow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	res := pow2(x, n>>1)
	res *= res
	if n&0x1 == 1 {
		res *= x
	}
	return res
}
