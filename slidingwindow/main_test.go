package main

import (
	"reflect"
	"testing"
)

func Test_slidingWindow(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get klm",
			args: args{
				str: "klm",
			},
			want: 3,
		},
		{
			name: "get kl",
			args: args{
				str: "kkllmm",
			},
			want: 2,
		},
		{
			name: "get kl",
			args: args{
				str: "kll",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingWindow(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("slidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
