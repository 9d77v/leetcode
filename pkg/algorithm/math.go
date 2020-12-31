package algorithm

//Max return the maximum number
func Max(a ...int) int {
	max := a[0]
	for _, v := range a[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

//Min return the minimum number
func Min(a ...int) int {
	min := a[0]
	for _, v := range a[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
