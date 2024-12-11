package interview

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		//["o","a","b","n"],["o","t","a","e"],["a","h","k","r"],["a","f","l","v"]
		{
			name: "case 1",
			args: args{
				board: [][]byte{
					{'o', 'a', 'b', 'n'},
					{'o', 't', 'a', 'e'},
					{'a', 'h', 'k', 'r'},
					{'a', 'f', 'l', 'v'},
				},
				words: []string{"oa", "oaa"},
			},
			want: []string{"oa", "oaa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
