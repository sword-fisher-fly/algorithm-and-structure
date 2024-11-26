package special

import "testing"

func TestLastStoneWeight(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				stones: []int{1, 1, 1, 1},
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				stones: []int{2, 7, 4, 1, 8, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastStoneWeight(tt.args.stones); got != tt.want {
				t.Errorf("LastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
