package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "four: [3,1,4,2]",
			input: []int{3, 1, 4, 2},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "ten: [3,6,8,10,4,2,1,5,7,9]",
			input: []int{3, 6, 8, 10, 4, 2, 1, 5, 7, 9},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
