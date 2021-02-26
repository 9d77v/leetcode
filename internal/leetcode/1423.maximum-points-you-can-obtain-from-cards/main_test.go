package main

import (
	"testing"
)

type args struct {
	cardPoints []int
	k          int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]int{1, 2, 3, 4, 5, 6, 1}, 3}, 12},
	{"should be equal", args{[]int{2, 2, 2}, 2}, 4},
	{"should be equal", args{[]int{9, 7, 7, 9, 7, 7, 9}, 7}, 55},
	{"should be equal", args{[]int{1, 1000, 1}, 1}, 1},
	{"should be equal", args{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3}, 202},
}

func Test_maxScore(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxScoreFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreFunc2(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScoreFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
