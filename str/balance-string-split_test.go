package str

import "testing"

func TestBalanceStringSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "RLRRLLRLRL",
			args: args{
				s: "RLRRLLRLRL",
			},
			want: 4,
		},
		{
			name: "RLLLLRRRLR",
			args: args{
				s: "RLLLLRRRLR",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BalanceStringSplit(tt.args.s); got != tt.want {
				t.Errorf("BalanceStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
