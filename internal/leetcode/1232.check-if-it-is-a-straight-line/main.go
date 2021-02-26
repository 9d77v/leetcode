package main

/*
题目：缀点成线
在一个 XY 坐标系中有一些点，我们用数组 coordinates 来分别记录它们的坐标，其中 coordinates[i] = [x, y] 表示横坐标为 x、纵坐标为 y 的点。

请你来判断，这些点是否在该坐标系中属于同一条直线上，是则返回 true，否则请返回 false。

提示：

2 <= coordinates.length <= 1000
coordinates[i].length == 2
-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
coordinates 中不含重复的点

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-if-it-is-a-straight-line
*/

func checkStraightLine(coordinates [][]int) bool {
	n := len(coordinates)
	x1, y1, x2, y2 := coordinates[0][0], coordinates[0][1], coordinates[n-1][0], coordinates[n-1][1]
	for i := 1; i < n-1; i++ {
		x, y := coordinates[i][0], coordinates[i][1]
		if (x-x1)*(y-y2) != (x-x2)*(y-y1) {
			return false
		}
	}
	return true
}
