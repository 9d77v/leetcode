package main

import (
	"testing"
)

type args struct {
	num int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{12258}, 5},
	{"should be equal", args{0}, 1},
}

func Test_translateNum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNum(tt.args.num); got != tt.want {
				t.Errorf("translateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateNumFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNumFunc2(tt.args.num); got != tt.want {
				t.Errorf("translateNumFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateNumFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNumFunc3(tt.args.num); got != tt.want {
				t.Errorf("translateNumFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
