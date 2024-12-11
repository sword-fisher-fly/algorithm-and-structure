package backtracking

import (
	"reflect"
	"testing"
)

func TestRemoveInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{s: "()())()"},
			want: []string{"(())()", "()()()"},
		},
		{
			name: "case 2",
			args: args{s: ")("},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveInvalidParentheses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveInvalidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
