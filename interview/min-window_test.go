package interview

import "testing"

func TestMinWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "s=ADOBECODEBANC,t=ABC",
			args: args{s: "ADOBECODEBANC", t: "ABC"},
			want: "BANC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("MinWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
