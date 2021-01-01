package main

import "testing"

type args struct {
	x float64
	n int
}

var tests = []struct {
	name string
	args args
	want float64
}{
	{"should be ok", args{2.0000, 10}, 1024},
	{"should be ok", args{2.1000, 3}, 9.261},
	{"should be ok", args{2.0000, -2}, 0.25},
}

func Test_myPowFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			precision := 1e-5
			if got := myPowFunc1(tt.args.x, tt.args.n); (got - tt.want) > precision {
				t.Errorf("myPowFunc1() = %v", (got - tt.want))
			}
		})
	}
}
