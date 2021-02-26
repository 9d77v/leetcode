package main

/*
题目：等价多米诺骨牌对的数量

给你一个由一些多米诺骨牌组成的列表 dominoes。

如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。

形式上，dominoes[i] = [a, b] 和 dominoes[j] = [c, d] 等价的前提是 a==c 且 b==d，或是 a==d 且 b==c。

在 0 <= i < j < dominoes.length 的前提下，找出满足 dominoes[i] 和 dominoes[j] 等价的骨牌对 (i, j) 的数量。

提示：

1 <= dominoes.length <= 40000
1 <= dominoes[i][j] <= 9


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/
*/

/*
方法一：记录访问过的数据
时间复杂度：О(n^2)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2 MB
*/
func numEquivDominoPairs(dominoes [][]int) (count int) {
	n := len(dominoes)
	if n == 0 {
		return 0
	}
	visited := make([]bool, len(dominoes))
	visited[0] = true
	for i := 0; i < n-1; {
		a, b := dominoes[i][0], dominoes[i][1]
		equalNum := 0
		for j := i + 1; j < n; j++ {
			if !visited[j] {
				c, d := dominoes[j][0], dominoes[j][1]
				if (a == c && b == d) || (a == d && b == c) {
					equalNum++
					visited[j] = true
				}
			}
		}
		if equalNum > 0 {
			count += equalNum * (equalNum + 1) / 2
		}
		k := i + 1
		for ; k < n; k++ {
			if !visited[k] {
				break
			}
		}
		i = k
	}
	return
}

/*
方法二：二元组表示+计数
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：32 ms	内存消耗：7.6 MB
*/
func numEquivDominoPairsFUnc2(dominoes [][]int) (count int) {
	nums := [100]int{}
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0]*10 + d[1]
		count += nums[v]
		nums[v]++
	}
	return
}
