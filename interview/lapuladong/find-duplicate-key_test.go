package lapuladong

import "testing"

func TestFindDuplicateNum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1,3,4,2,2",
			args: args{
				nums: []int{1, 3, 4, 2, 2},
			},
			want: 2,
		},
		{
			name: "3,1,3,4,2",
			args: args{
				nums: []int{3, 1, 3, 4, 2},
			},
			want: 3,
		},
		{
			name: "3,3,3,3,3",
			args: args{
				nums: []int{3, 3, 3, 3, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicateNum(tt.args.nums); got != tt.want {
				t.Errorf("FindDuplicateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
