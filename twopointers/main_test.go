package main

import (
	"reflect"
	"testing"
)

func Test_twoPointers(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "target 9",
			args: args{
				nums:   []int{2, 7},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "target 2, return empty",
			args: args{
				nums:   []int{3, 7},
				target: 2,
			},
			want: []int{},
		},
		{
			name: "target 12",
			args: args{
				nums:   []int{6, 2, 3, 9, 10, 4, 4},
				target: 12,
			},
			want: []int{0, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoPointers(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoPointers() = %v, want %v", got, tt.want)
			}
		})
	}
}
