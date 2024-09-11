package stack

import "testing"

func TestEvalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1+1=2",
			args: args{tokens: []string{"1", "1", "+"}},
			want: 2,
		},
		{
			name: "1+2=3",
			args: args{tokens: []string{"2", "1", "+"}},
			want: 3,
		},
		{
			name: "(1+2)*3=9",
			args: args{tokens: []string{"2", "1", "+", "3", "*"}},
			want: 9,
		},
		{
			name: "14/5+4=6",
			args: args{tokens: []string{"14", "5", "/", "4", "+"}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("EvalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
