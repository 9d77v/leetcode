package main

import "sort"

/*
题目：
你打算做甜点，现在需要购买配料。目前共有 n 种冰激凌基料和 m 种配料可供选购。而制作甜点需要遵循以下几条规则：

必须选择 一种 冰激凌基料。
可以添加 一种或多种 配料，也可以不添加任何配料。
每种类型的配料 最多两份 。
给你以下三个输入：

baseCosts ，一个长度为 n 的整数数组，其中每个 baseCosts[i] 表示第 i 种冰激凌基料的价格。
toppingCosts，一个长度为 m 的整数数组，其中每个 toppingCosts[i] 表示 一份 第 i 种冰激凌配料的价格。
target ，一个整数，表示你制作甜点的目标价格。
你希望自己做的甜点总成本尽可能接近目标价格 target 。

返回最接近 target 的甜点成本。如果有多种方案，返回 成本相对较低 的一种。

提示：

n == baseCosts.length
m == toppingCosts.length
1 <= n, m <= 10
1 <= baseCosts[i], toppingCosts[i] <= 10^4
1 <= target <= 10^4
*/

func closestCost(baseCosts []int, toppingCosts []int, target int) (max int) {
	sort.Ints(baseCosts)
	sort.Ints(toppingCosts)
	for _, baseCost := range baseCosts {
		if baseCost > target {
			break
		}
		for _, toppingCost := range toppingCosts {
			if baseCost+2*toppingCost <= target {
				max = Max(max, baseCost+2*toppingCost)
			} else if baseCost+toppingCost <= target {
				max = Max(max, baseCost+toppingCost)
			}
		}
		if max == target {
			return max
		}
	}
	if max == 0 {
		return baseCosts[0]
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
