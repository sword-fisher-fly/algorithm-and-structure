package special

import "testing"

func TestNumRollsToTarget(t *testing.T) {
	type args struct {
		dices  int
		faces  int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{1, 6, 3},
			want: 1,
		},
		{
			name: "case 2",
			args: args{2, 6, 7},
			want: 6,
		},
		{
			name: "case 3",
			args: args{30, 30, 500},
			want: 222616187,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumRollsToTarget(tt.args.dices, tt.args.faces, tt.args.target); got != tt.want {
				t.Errorf("NumRollsToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
