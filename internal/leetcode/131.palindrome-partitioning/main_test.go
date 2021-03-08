package main

import (
	"reflect"
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want [][]string
}{
	{"should be equal", args{"aab"}, [][]string{
		{"a", "a", "b"}, {"aa", "b"},
	}},
}

func Test_partition(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := partition(tt.args.s); !reflect.DeepEqual(gotAns, tt.want) {
				t.Errorf("partition() = %v, want %v", gotAns, tt.want)
			}
		})
	}
}

func Test_partitionFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionFunc2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partitionFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionFunc3(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
