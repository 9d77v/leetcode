package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/unionfind"
)

/*
题目：等式方程的可满足性
给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。

只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。

提示：
1 <= equations.length <= 500
equations[i].length == 4
equations[i][0] 和 equations[i][3] 是小写字母
equations[i][1] 要么是 '='，要么是 '!'
equations[i][2] 是 '='

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/satisfiability-of-equality-equations
*/

/*
方法一：并查集(数组)
时间复杂度：О(n+ClogC)
空间复杂度：О(C)
运行时间：4 ms	内存消耗：2.8 MB
*/
func equationsPossible(equations []string) bool {
	var uf UnionFind = NewArrayUnionFind(26)
	for _, equation := range equations {
		a, b := int(equation[0]-'a'), int(equation[3]-'a')
		if equation[1] == '=' {
			uf.Union(a, b)
		}
	}
	for _, equation := range equations {
		if equation[1] == '!' {
			a, b := int(equation[0]-'a'), int(equation[3]-'a')
			uf.Union(a, a)
			uf.Union(b, b)
			if uf.IsConnected(a, b) {
				return false
			}
		}
	}
	return true
}
