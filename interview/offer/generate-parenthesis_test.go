package offer

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1 pair bracket",
			args: args{n: 1},
			want: []string{"()"},
		},
		{
			name: "2 pair bracket",
			args: args{n: 2},
			want: []string{"(())", "()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateParenthesisII(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1 pair bracket",
			args: args{n: 1},
			want: []string{"()"},
		},
		{
			name: "2 pair bracket",
			args: args{n: 2},
			want: []string{"(())", "()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateParenthesisII(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateParenthesisII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateParenthesisIII(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1 pair bracket",
			args: args{n: 1},
			want: []string{"()"},
		},
		{
			name: "2 pair bracket",
			args: args{n: 2},
			want: []string{"()()", "(())"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateParenthesisIII(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateParenthesisIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
