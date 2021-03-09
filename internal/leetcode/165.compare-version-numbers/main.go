package main

import (
	"strconv"
	"strings"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目： 比较版本号
给你两个版本号 version1 和 version2 ，请你比较它们。

版本号由一个或多个修订号组成，各修订号由一个 '.' 连接。每个修订号由 多位数字 组成，可能包含 前导零 。每个版本号至少包含一个字符。修订号从左到右编号，下标从 0 开始，最左边的修订号下标为 0 ，下一个修订号下标为 1 ，以此类推。例如，2.5.33 和 0.1 都是有效的版本号。

比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较 忽略任何前导零后的整数值 。也就是说，修订号 1 和修订号 001 相等 。如果版本号没有指定某个下标处的修订号，则该修订号视为 0 。例如，版本 1.0 小于版本 1.1 ，因为它们下标为 0 的修订号相同，而下标为 1 的修订号分别为 0 和 1 ，0 < 1 。

返回规则如下：

如果 version1 > version2 返回 1，
如果 version1 < version2 返回 -1，
除此之外返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/compare-version-numbers
*/

/*
方法一：分割+解析
时间复杂度：О(n+m+max(m,n))
空间复杂度：О(n+m)
运行时间：0 ms	内存消耗：1.9 MB
*/
func compareVersion(version1 string, version2 string) int {
	arr1, arr2 := strings.Split(version1, "."), strings.Split(version2, ".")
	for i := 0; i < Max(len(arr1), len(arr2)); i++ {
		num1, num2 := 0, 0
		if i < len(arr1) {
			num1, _ = strconv.Atoi(arr1[i])
		}
		if i < len(arr2) {
			num2, _ = strconv.Atoi(arr2[i])
		}
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
	}
	return 0
}
