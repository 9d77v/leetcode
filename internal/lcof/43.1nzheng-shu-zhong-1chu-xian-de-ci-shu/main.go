package main

/*
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

限制：

1 <= n < 2^31

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof
*/

/*
方法一: 找数学规律
时间复杂度：О(logN)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：1.9 MB
*/
func countDigitOne(n int) int {
	high, cur, low, digit, count := n/10, n%10, 0, 1, 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			count += high * digit
		} else if cur == 1 {
			count += high*digit + low + 1
		} else {
			count += (high + 1) * digit
		}
		high, cur, low, digit = high/10, high%10, low+cur*digit, digit*10
	}
	return count
}
