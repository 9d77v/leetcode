package main

import (
	"testing"
)

type args struct {
	arr []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}}, 5},
	{"should be equal", args{[]int{4, 8, 12, 16}}, 2},
	{"should be equal", args{[]int{100}}, 1},
	{"should be equal", args{[]int{9, 9}}, 1},
	{"should be equal", args{[]int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0}}, 5},
	{"should be equal", args{[]int{2, 0, 2, 4, 2, 5, 0, 1, 2, 3}}, 6},
}

func Test_maxTurbulenceSize(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSize(tt.args.arr); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxTurbulenceSizeFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSizeFunc2(tt.args.arr); got != tt.want {
				t.Errorf("maxTurbulenceSizeFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxTurbulenceSizeFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSizeFunc3(tt.args.arr); got != tt.want {
				t.Errorf("maxTurbulenceSizeFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
