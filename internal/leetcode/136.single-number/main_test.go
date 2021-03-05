package main

import "testing"

type args struct {
	nums []int
}

var tests = []struct {
	name       string
	args       args
	wantResult int
}{
	{"should be equal", args{[]int{2, 2, 1}}, 1},
	{"should be equal", args{[]int{4, 1, 2, 1, 2}}, 4},
}

func Test_singleNumber(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := singleNumber(tt.args.nums); gotResult != tt.wantResult {
				t.Errorf("singleNumber() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
