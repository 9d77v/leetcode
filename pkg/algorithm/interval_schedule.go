package algorithm

import "sort"

//OverlapIntervalSchedule 边界重叠的区间调度
func OverlapIntervalSchedule(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res, right := 1, intervals[0][1]
	for _, v := range intervals[1:] {
		if v[0] >= right {
			res++
			right = v[1]
		}
	}
	return res
}

//IntervalSchedule 区间调度
func IntervalSchedule(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res, right := 1, intervals[0][1]
	for _, v := range intervals[1:] {
		if v[0] > right {
			res++
			right = v[1]
		}
	}
	return res
}
