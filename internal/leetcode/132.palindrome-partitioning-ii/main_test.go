package main

import "testing"

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{"aab"}, 1},
	{"should be equal", args{"aa"}, 0},
	{"should be equal", args{"ab"}, 1},
	{"should be equal", args{"ccaacabacb"}, 3},
	{"should be equal", args{"eegiicgaeadbcfacfhifdbiehbgejcaeggcgbahfcajfhjjdgj"}, 42},
}

func Test_minCut(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
