package str

import (
	"testing"
)

func TestFindStr(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{s: "dabcd", t: "a[bc]"},
			want: 1,
		},
		{
			name: "case 2",
			args: args{s: "dabcd", t: "b[cd]"},
			want: 2,
		},
		{
			name: "case 1",
			args: args{s: "dabcd", t: "a[bc]"},
			want: 1,
		},
		{
			name: "case 3",
			args: args{s: "dabcd", t: "ef"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindStr(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("FindStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestFindStrII(t *testing.T) {
// 	type args struct {
// 		s string
// 		t string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			name: "case 1",
// 			args: args{s: "dabcd", t: "a[bc]"},
// 			want: 1,
// 		},
// 		{
// 			name: "case 2",
// 			args: args{s: "dabcd", t: "b[cd]"},
// 			want: 2,
// 		},
// 		{
// 			name: "case 1",
// 			args: args{s: "dabcd", t: "a[bc]"},
// 			want: 1,
// 		},
// 		{
// 			name: "case 3",
// 			args: args{s: "dabcd", t: "ef"},
// 			want: -1,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := FindStrII(tt.args.s, tt.args.t); got != tt.want {
// 				t.Errorf("FindStrII() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
