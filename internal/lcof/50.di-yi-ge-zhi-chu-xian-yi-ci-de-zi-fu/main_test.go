package main

import "testing"

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want byte
}{
	{"should be equal", args{"abaccdeff"}, 'b'},
	{"should be equal", args{""}, ' '},
}

func Test_firstUniqChar(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
