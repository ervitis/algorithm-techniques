package main

import "testing"

func Test_graph_dfs(t *testing.T) {
	type args struct {
		start   rune
		target  rune
		visited map[rune]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A -> F",
			args: args{
				start:   'A',
				target:  'F',
				visited: make(map[rune]bool),
			},
			want: true,
		},
		{
			name: "A -> I",
			args: args{
				start:   'A',
				target:  'I',
				visited: make(map[rune]bool),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGraph()
			g.addEdge('A', 'C')
			g.addEdge('C', 'E')
			g.addEdge('E', 'G')
			g.addEdge('G', 'F')
			g.addEdge('A', 'D')
			g.addEdge('D', 'F')
			g.addEdge('E', 'B')
			g.addEdge('B', 'H')
			g.addEdge('I', 'H')
			if got := g.dfs(tt.args.start, tt.args.target, tt.args.visited); got != tt.want {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
