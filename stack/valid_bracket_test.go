package stack

import "testing"

func TestIsValidBracket(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty string",
			args: args{},
			want: true,
		},
		{
			name: "valid string 1",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "valid string 2",
			args: args{
				s: "(([][]{}{}))",
			},
			want: true,
		},
		{
			name: "invalid string 1",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "invalid string 2",
			args: args{
				s: "([)]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBracket(tt.args.s); got != tt.want {
				t.Errorf("IsValidBracket() = %v, want %v", got, tt.want)
			}
		})
	}
}
