package algorithm

import "math"

//Max return the maximum number
func Max(a ...int) int {
	if len(a) == 0 {
		return math.MinInt64
	}
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

//Min return the minimum number
func Min(a ...int) int {
	if len(a) == 0 {
		return math.MinInt64
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
