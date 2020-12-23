package main

import (
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test leetcode", args{"leetcode"}, 0},
	{"test loveleetcode", args{"loveleetcode"}, 2},
	{"test ''", args{""}, -1},
}

func Test_firstUniqCharFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqCharFunc1(tt.args.s); got != tt.want {
				t.Errorf("firstUniqCharFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstUniqCharFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqCharFunc2(tt.args.s); got != tt.want {
				t.Errorf("firstUniqCharFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstUniqCharFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqCharFunc3(tt.args.s); got != tt.want {
				t.Errorf("firstUniqCharFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
