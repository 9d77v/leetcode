package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：数组形式的整数加法
对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。

给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。

提示：

1 <= A.length <= 10000
0 <= A[i] <= 9
0 <= K <= 10000
如果 A.length > 1，那么 A[0] != 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-to-array-form-of-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
方法一：逐位相加
时间复杂度：О(max(n,logK))
空间复杂度：О(max(n,logK))
运行时间：36 ms	内存消耗：6.7 MB
*/
func addToArrayForm(A []int, K int) (result []int) {
	i, carry := len(A)-1, 0
	for K > 0 || i >= 0 || carry == 1 {
		kNum := K % 10
		K /= 10
		aNum := 0
		if i >= 0 {
			aNum = A[i]
		}
		sum := aNum + kNum + carry
		carry = sum / 10
		i--
		result = append(result, sum%10)
	}
	Reverse(result)
	return
}
