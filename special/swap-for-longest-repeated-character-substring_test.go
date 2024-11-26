package special

import "testing"

func TestMaxRepOpt1(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				text: "ababa",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				text: "aaabaaa",
			},
			want: 6,
		},
		{
			name: "case 3",
			args: args{
				text: "aaabbbbbaaa",
			},
			want: 5,
		},
		{
			name: "case 4",
			args: args{
				text: "aaabbaaa",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxRepOpt1(tt.args.text); got != tt.want {
				t.Errorf("MaxRepOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
