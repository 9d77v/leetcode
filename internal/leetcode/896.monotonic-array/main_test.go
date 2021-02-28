package main

import (
	"testing"
)

type args struct {
	A []int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be true", args{[]int{1, 2, 2, 3}}, true},
	{"should be true", args{[]int{6, 5, 4, 4}}, true},
	{"should be false", args{[]int{1, 3, 2}}, false},
	{"should be true", args{[]int{1, 2, 4, 5}}, true},
	{"should be true", args{[]int{1, 1, 1}}, true},
}

func Test_isMonotonic(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.A); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMonotonicFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonicFunc2(tt.args.A); got != tt.want {
				t.Errorf("isMonotonicFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMonotonicFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonicFunc3(tt.args.A); got != tt.want {
				t.Errorf("isMonotonicFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMonotonicFunc4(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonicFunc4(tt.args.A); got != tt.want {
				t.Errorf("isMonotonicFunc4() = %v, want %v", got, tt.want)
			}
		})
	}
}
