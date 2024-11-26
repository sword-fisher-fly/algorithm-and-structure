package special

import "testing"

func TestAlphabetBoardPath(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				target: "dfdsfkjdfleiwllsgoidkjflsdfjdfierwsd",
			},
			want: "RRR!LLLD!URRR!DDD!LLLUU!D!URRRR!LU!LLLD!DR!UURRR!LD!LDDD!LUU!!DRR!LLUU!DRRR!LU!U!LLLDD!URRRR!LLLL!DR!DRR!UUU!LLLD!RRRR!LU!LLLD!RRR!UR!LLDDD!D!UR!UUU!",
		},
		{
			name: "case 2",
			args: args{
				target: "zdz",
			},
			want: "DDDDD!UUUUURRR!LLLDDDDD!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlphabetBoardPath(tt.args.target); got != tt.want {
				t.Errorf("AlphabetBoardPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
