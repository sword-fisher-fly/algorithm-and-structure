package tree

import (
	"testing"
)

func TestIsSame(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "the two tree is same",
			args: args{
				p: NewTreeFromArray(level4Arr).Root(),
				q: NewTreeFromArray(level4Arr).Root(),
			},
			want: true,
		},
		{
			name: "the two tree is not same",
			args: args{
				p: NewTreeFromArray(level4Arr).Root(),
				q: NewTreeFromArray(level3Arr).Root(),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSame(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("IsSame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSameByBFS(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "the two tree is same",
			args: args{
				p: NewTreeFromArray(level4Arr).Root(),
				q: NewTreeFromArray(level4Arr).Root(),
			},
			want: true,
		},
		{
			name: "the two tree is not same",
			args: args{
				p: NewTreeFromArray(level4Arr).Root(),
				q: NewTreeFromArray(level3Arr).Root(),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameByBFS(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("IsSameByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
