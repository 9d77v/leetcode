package main

import (
	"testing"
)

type args struct {
	A []int
	K int
}

var tests = []struct {
	name       string
	args       args
	wantResult int
}{
	{"should be equal", args{[]int{0, 1, 0}, 1}, 2},
	{"should be equal", args{[]int{1, 1, 0}, 2}, -1},
	{"should be equal", args{[]int{0, 1, 1}, 2}, -1},
	{"should be equal", args{[]int{0, 0, 0, 1, 0, 1, 1, 0}, 3}, 3},
}

func Test_minKBitFlips(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minKBitFlips(tt.args.A, tt.args.K); got != tt.wantResult {
				t.Errorf("minKBitFlips() = %v, wantResult %v", got, tt.wantResult)
			}
		})
	}
}

func Test_minKBitFlipsFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := minKBitFlipsFunc2(tt.args.A, tt.args.K); gotResult != tt.wantResult {
				t.Errorf("minKBitFlipsFunc2() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_minKBitFlipsFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := minKBitFlipsFunc3(tt.args.A, tt.args.K); gotResult != tt.wantResult {
				t.Errorf("minKBitFlipsFunc3() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
