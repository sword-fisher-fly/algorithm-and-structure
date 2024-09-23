package str

import "testing"

// 1)  0
// 2) +0
// 3) -0
// 4) 1.0
// 5) -1.0
// 6) +1.0
func Test_isValidRealNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0",
			args: args{s: "0"},
			want: true,
		},
		{
			name: "+0",
			args: args{s: "+0"},
			want: true,
		},
		{
			name: "-0",
			args: args{s: "-0"},
			want: true,
		},
		{
			name: "+1.0",
			args: args{s: "+1.0"},
			want: true,
		},
		{
			name: "1.0",
			args: args{s: "1.0"},
			want: true,
		},
		{
			name: "-1.0",
			args: args{s: "-1.0"},
			want: true,
		},
		{
			name: "two many sign",
			args: args{s: "--1.0"},
			want: false,
		},
		{
			name: "two many leading zero",
			args: args{s: "00"},
			want: false,
		},
		{
			name: "two many dot",
			args: args{s: "0..1"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidRealNumber(tt.args.s); got != tt.want {
				t.Errorf("isValidRealNumber(%s) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
