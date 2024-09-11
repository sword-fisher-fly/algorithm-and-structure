package greedy

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 stair",
			args: args{cost: []int{1}},
			want: 0,
		},
		{
			name: "2 stairs",
			args: args{cost: []int{10, 15}},
			want: 0,
		},
		{
			name: "3 stairs",
			args: args{cost: []int{10, 15, 20}},
			want: 15,
		},
		{
			name: "10 stairs",
			args: args{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("MinCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinCostClimbingStairsOptimize(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 stair",
			args: args{cost: []int{1}},
			want: 0,
		},
		{
			name: "2 stairs",
			args: args{cost: []int{10, 15}},
			want: 0,
		},
		{
			name: "3 stairs",
			args: args{cost: []int{10, 15, 20}},
			want: 15,
		},
		{
			name: "10 stairs",
			args: args{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCostClimbingStairsOptimize(tt.args.cost); got != tt.want {
				t.Errorf("MinCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
