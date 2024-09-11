package greedy

import "testing"

// 问题: 一条环形公路，有N个加油站，gas[i]表示第i个加油站的汽油量，cost[i]表示从第i个加油站到第i+1个加油站的消耗量。
// 输出: 如从能够到达第N个加油站，返回出发加油站的编号，否则返回-1。
// Example:
// gas: [1,2,3,4,5]
// cost:[3,4,5,1,2]
// output: 3

func TestCanCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero station",
			args: args{
				gas:  []int{},
				cost: []int{},
			},
			want: -1,
		},
		{
			name: "1 station",
			args: args{
				gas:  []int{1},
				cost: []int{1},
			},
			want: 0,
		},
		{
			name: "5 station",
			args: args{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("CanCompleteCircuit() = %v, want %v", got, tt.want)
			}

			if got := CanCompleteCircuitII(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("CanCompleteCircuitII() = %v, want %v", got, tt.want)
			}
		})
	}
}
