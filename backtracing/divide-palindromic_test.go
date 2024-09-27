package backtracking

import (
	"reflect"
	"testing"
)

func TestPalindromePartition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "two palindrome partitions",
			args: args{s: "aab"},
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "",
			args: args{s: "abaccda"},
			want: [][]string{
				{"a", "b", "a", "c", "c", "d", "a"},
				{"a", "b", "a", "cc", "d", "a"},
				{"aba", "c", "c", "d", "a"},
				{"aba", "cc", "d", "a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromePartition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PalindromePartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPalindromePartitionII(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "two palindrome partitions",
			args: args{s: "aab"},
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "",
			args: args{s: "abaccda"},
			want: [][]string{
				{"a", "b", "a", "c", "c", "d", "a"},
				{"a", "b", "a", "cc", "d", "a"},
				{"aba", "c", "c", "d", "a"},
				{"aba", "cc", "d", "a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromePartitionII(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PalindromePartitionII() = %v, want %v", got, tt.want)
			}
		})
	}
}
