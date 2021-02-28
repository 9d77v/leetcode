package main

import "testing"

type args struct {
	baseCosts    []int
	toppingCosts []int
	target       int
}

var tests = []struct {
	name string
	args args
	want int
}{
	// {"should be equal", args{[]int{1, 7}, []int{3, 4}, 10}, 10},
	{"should be equal", args{[]int{2, 3}, []int{4, 5, 100}, 18}, 17},
	// {"should be equal", args{[]int{3, 10}, []int{2, 5}, 9}, 8},
	// {"should be equal", args{[]int{10}, []int{1}, 1}, 10},
}

func Test_closestCost(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestCost(tt.args.baseCosts, tt.args.toppingCosts, tt.args.target); got != tt.want {
				t.Errorf("closestCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
