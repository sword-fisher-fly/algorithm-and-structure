package hash

import "testing"

func TestFourSumCount(t *testing.T) {
	type args struct {
		A []int64
		B []int64
		C []int64
		D []int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "has 2 results",
			// A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
			// A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
			args: args{
				A: []int64{1, 2},
				B: []int64{-2, -1},
				C: []int64{-1, 2},
				D: []int64{0, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FourSumCount(tt.args.A, tt.args.B, tt.args.C, tt.args.D); got != tt.want {
				t.Errorf("FourSumCount() = %v, want %v", got, tt.want)
			}

			if got := FourSumCountII(tt.args.A, tt.args.B, tt.args.C, tt.args.D); got != tt.want {
				t.Errorf("FourSumCountII() = %v, want %v", got, tt.want)
			}
		})
	}
}
