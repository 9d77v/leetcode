package main

import (
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name    string
	args    args
	wantMax int
}{
	{"should be equal", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
	{"should be equal", args{[]int{0, 1, 0, 3, 2, 3}}, 4},
	{"should be equal", args{[]int{7, 7, 7, 7, 7, 7}}, 1},
}

func Test_lengthOfLIS(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := lengthOfLIS(tt.args.nums); gotMax != tt.wantMax {
				t.Errorf("lengthOfLIS() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func Test_lengthOfLISFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := lengthOfLISFunc2(tt.args.nums); gotMax != tt.wantMax {
				t.Errorf("lengthOfLISFunc2() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
