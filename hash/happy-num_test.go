package hash

import "testing"

func TestIsHappy(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy 100",
			args: args{n: 100},
			want: true,
		},
		{
			name: "12345",
			args: args{n: 12345},
			want: false,
		},
		{
			name: "happy 19",
			args: args{n: 19},
			want: true,
		},

		{
			name: "not happy 2",
			args: args{n: 2},
			want: false,
		},
		{
			name: "not happy 4",
			args: args{n: 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHappy(tt.args.n); got != tt.want {
				t.Errorf("IsHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
