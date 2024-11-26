package huawei

import "testing"

// https://www.nowcoder.com/practice/f9a4c19050fc477e9e27eb75f3bfd49c?tpId=37&tqId=21264&rp=1&ru=/exam/oj/ta&qru=/exam/oj/ta&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D37&difficulty=undefined&judgeStatus=undefined&tags=&title=

func TestTotalDifferentWeights(t *testing.T) {
	type args struct {
		weights []int
		nums    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				weights: []int{1, 2},
				nums:    []int{2, 1},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalDifferentWeights(tt.args.weights, tt.args.nums); got != tt.want {
				t.Errorf("TotalDifferentWeights() = %v, want %v", got, tt.want)
			}
		})
	}
}
