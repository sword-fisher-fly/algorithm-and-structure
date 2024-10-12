package offer

import "testing"

func TestIsBipartite(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "non-bipartite",
			args: args{graph: [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}},
			want: false,
		},
		{
			name: "bipartite",
			args: args{graph: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBipartite(tt.args.graph); got != tt.want {
				t.Errorf("IsBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
