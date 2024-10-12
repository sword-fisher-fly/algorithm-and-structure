package interview

import "testing"

func TestEvalExpr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2+3-7",
			args: args{
				s: "1+2+3-7",
			},
			want: -1,
		},
		{
			name: "1+2*3-8/4",
			args: args{
				s: "1+2*3-8/4",
			},
			want: 5,
		},
		{
			name: "(1+3)*4-(6-8)",
			args: args{
				s: "(1+3)*4-(6-8)",
			},
			want: 18,
		},
		{
			name: "(1+3)* 4-(6-8) with space character",
			args: args{
				s: "(1+3)* 4-(6-8)",
			},
			want: 18,
		},
		{
			name: "(1+3)*4-3%4+3",
			args: args{
				s: "(1+3)*4-3%4+3",
			},
			want: 16,
		},
		{
			name: "(1+3)*4-3*4+3",
			args: args{
				s: "(1+3)*4-3*4+3",
			},
			want: 7,
		},
		{
			name: "(1+3)*4-3*((7+3)/2+1)",
			args: args{
				s: "(1+3)*4-3*((7+3)/2+1)",
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalExpr(tt.args.s); got != tt.want {
				t.Errorf("EvalExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
