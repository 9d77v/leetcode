package main

import (
	"testing"
)

type args struct {
	envelopes [][]int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[][]int{
		{5, 4}, {6, 4}, {6, 7}, {2, 3},
	}}, 3},
	{"should be equal", args{[][]int{
		{2, 100}, {3, 200}, {4, 300}, {5, 500}, {5, 400}, {5, 250}, {6, 370}, {6, 360}, {7, 380},
	}}, 5},
}

func Test_maxEnvelopes(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxEnvelopesFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxEnvelopesFunc2(tt.args.envelopes); gotMax != tt.want {
				t.Errorf("maxEnvelopesFunc2() = %v, want %v", gotMax, tt.want)
			}
		})
	}
}
