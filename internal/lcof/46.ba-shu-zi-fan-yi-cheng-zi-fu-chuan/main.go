package main

import "strconv"

/*
题目：
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

提示：

0 <= num < 2^31

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
*/

/*
方法一：动态规划 f(0)=f(1)=1 f(i)=f(i-1) nums[i-1]nums[i]>26 f(i)=f(i-2)+f(i-1) nums[i-1]nums[i]<26
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：1.9 MB
*/
func translateNum(num int) int {
	s := strconv.Itoa(num)
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < len(s)+1; i++ {
		tmp, _ := strconv.Atoi(s[i-2 : i])
		if tmp >= 10 && tmp <= 25 {
			dp[i] = dp[i-2] + dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(s)]
}

/*
方法二：动态规划 优化 f(0)=f(1)=1 f(i)=f(i-1) nums[i-1]nums[i]>26 f(i)=f(i-2)+f(i-1) nums[i-1]nums[i]<26
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：1.9 MB
*/
func translateNumFunc2(num int) int {
	s := strconv.Itoa(num)
	p, q, r := 1, 1, 1
	for i := 2; i < len(s)+1; i++ {
		tmp, _ := strconv.Atoi(s[i-2 : i])
		if tmp >= 10 && tmp <= 25 {
			r = p + q
		} else {
			r = q
		}
		p, q = q, r
	}
	return r
}

/*
方法二：数字求余
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：1.9 MB
*/
func translateNumFunc3(num int) int {
	p, q, r, x, y := 1, 1, 1, num%10, num%10
	for num != 0 {
		num /= 10
		x = num % 10
		tmp := 10*x + y
		if tmp >= 10 && tmp <= 25 {
			r = p + q
		} else {
			r = q
		}
		y, p, q = x, q, r
	}
	return r
}
