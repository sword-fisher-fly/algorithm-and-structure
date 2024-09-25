package greedy

import (
	"reflect"
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	type args struct {
		peoples [][2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		//[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
		//[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
		{
			name: "rearrange the people's height in the specified order",
			args: args{
				peoples: [][2]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			},
			want: [][2]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReconstructQueue(tt.args.peoples); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReconstructQueueII(t *testing.T) {
	type args struct {
		peoples [][2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		//[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
		//[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
		{
			name: "rearrange the people's height in the specified order",
			args: args{
				peoples: [][2]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			},
			want: [][2]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReconstructQueueII(tt.args.peoples); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReconstructQueueII() = %v, want %v", got, tt.want)
			}
		})
	}
}
