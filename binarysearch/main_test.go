package main

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find 4",
			args: args{
				nums:   []int{2, 3, 4, 8, 9, 10, 14},
				target: 4,
			},
			want: 2,
		},
		{
			name: "not found",
			args: args{
				nums:   []int{2, 3, 4, 8, 9, 10, 14},
				target: 7,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
