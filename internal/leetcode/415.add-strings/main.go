package main

/*
题目：字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

提示：

num1 和num2 的长度都小于 5100
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
*/

/*
方法一：逐位相加
时间复杂度：О(max(m,n))
空间复杂度：О(max(m,n))
运行时间：0 ms	内存消耗：2.5 MB
*/
func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	var carry byte = 0
	result := make([]byte, 0)
	for i >= 0 || j >= 0 || carry != 0 {
		var x byte = 0
		if i >= 0 {
			x = num1[i] - '0'
			i--
		}
		var y byte = 0
		if j >= 0 {
			y = num2[j] - '0'
			j--
		}
		sum := x + y + carry
		carry = sum / 10
		result = append(result, sum%10+'0')
	}
	reverse(result)
	return string(result)
}

func reverse(a []byte) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
