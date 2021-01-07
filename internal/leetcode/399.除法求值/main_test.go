package main

import (
	"reflect"
	"testing"
)

type args struct {
	equations [][]string
	values    []float64
	queries   [][]string
}

var tests = []struct {
	name string
	args args
	want []float64
}{
	{"should be ok", args{
		[][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0},
		[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
	}, []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}},
	{"should be ok", args{
		[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0},
		[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
	}, []float64{3.75000, 0.40000, 5.00000, 0.20000}},
	{"should be ok", args{
		[][]string{{"a", "b"}}, []float64{0.5},
		[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
	}, []float64{0.5000, 2.00000, -1.00000, -1.0000}},
	{"should be ok", args{
		[][]string{{"a", "b"}, {"c", "d"}}, []float64{1.0, 1.0},
		[][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}},
	}, []float64{-1.000, -1.00000, 1.00000, 1.0000}},
}

func Test_calcEquationFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquationFunc1(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquationFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcEquationFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquationFunc2(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquationFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_calcEquationFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquationFunc3(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquationFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
