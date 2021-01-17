package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/topsort"
)

/*
题目：
你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]

给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/course-schedule
*/

/*
方法一：拓扑排序(广度优先搜索)
时间复杂度：O(n+m)
空间复杂度：O(n+m)
运行时间：12 ms	内存消耗：6 MB
*/
func canFinishFunc1(numCourses int, prerequisites [][]int) bool {
	BfsTopSort(numCourses, prerequisites, func(u int) {
		numCourses--
	})
	return numCourses == 0
}

/*
方法二：拓扑排序(深度优先搜索)
时间复杂度：O(n+m)
空间复杂度：O(n+m)
运行时间：20 ms	内存消耗：6.1 MB
*/
func canFinishFunc2(numCourses int, prerequisites [][]int) bool {
	DfsTopSort(numCourses, prerequisites, true, func(u int) {
		numCourses--
	})
	return numCourses == 0
}
