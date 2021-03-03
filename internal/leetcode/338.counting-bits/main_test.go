package main

import (
	"reflect"
	"testing"
)

type args struct {
	num int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{2}, []int{0, 1, 1}},
	{"should be equal", args{5}, []int{0, 1, 1, 2, 1, 2}},
}

func Test_countBits(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBitsFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsFunc2(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBitsFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsFunc3(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBitsFunc4(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsFunc4(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsFunc4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBitsFunc5(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsFunc5(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsFunc5() = %v, want %v", got, tt.want)
			}
		})
	}
}
