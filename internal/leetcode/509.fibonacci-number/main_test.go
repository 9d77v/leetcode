package main

import (
	"testing"
)

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test f(2)", args{2}, 1},
	{"test f(3)", args{3}, 2},
	{"test f(4)", args{4}, 3},
}

func Test_fibFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibFunc1(tt.args.n); got != tt.want {
				t.Errorf("fibFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibFunc2(tt.args.n); got != tt.want {
				t.Errorf("fibFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibFunc3(tt.args.n); got != tt.want {
				t.Errorf("fibFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibFunc4(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibFunc4(tt.args.n); got != tt.want {
				t.Errorf("fibFunc4() = %v, want %v", got, tt.want)
			}
		})
	}
}
