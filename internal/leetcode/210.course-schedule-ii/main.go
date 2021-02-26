package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/topsort"
)

/*
题目：课程表II
现在你总共有 n 门课需要选，记为 0 到 n-1。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]

给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。

可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

说明:
输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。

提示:
这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
拓扑排序也可以通过 BFS 完成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/course-schedule-ii
*/

/*
方法一：拓扑排序(广度优先搜索)
时间复杂度：O(n+m)
空间复杂度：O(n+m)
运行时间：8 ms	内存消耗：6.2 MB
*/
func findOrderFunc1(numCourses int, prerequisites [][]int) []int {
	result := make([]int, 0)
	BfsTopSort(numCourses, prerequisites, func(u int) {
		result = append(result, u)
	})
	if len(result) == numCourses {
		return result
	}
	return []int{}
}

/*
方法二：拓扑排序(深度优先搜索)
时间复杂度：O(n+m)
空间复杂度：O(n+m)
运行时间：20 ms	内存消耗：6.3 MB
*/
func findOrderFunc2(numCourses int, prerequisites [][]int) []int {
	result := make([]int, 0)
	valid := true
	DfsTopSort(numCourses, prerequisites, valid, func(u int) {
		result = append(result, u)
	})
	if !valid {
		return []int{}
	}
	Reverse(result)
	return result
}
