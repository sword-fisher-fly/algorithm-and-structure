package offer

import (
	"testing"
)

func TestReplaceWord(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s:          "the cattle was rattled by the battery",
				dictionary: []string{"cat", "bat", "rat"},
			},
			want: "the cat was rat by the bat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceWord(tt.args.s, tt.args.dictionary); got != tt.want {
				t.Errorf("ReplaceWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceWordII(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s:          "the cattle was rattled by the battery",
				dictionary: []string{"cat", "bat", "rat"},
			},
			want: "the cat was rat by the bat",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceWordII(tt.args.s, tt.args.dictionary); got != tt.want {
				t.Errorf("ReplaceWordII() = %v, want %v", got, tt.want)
			}
		})
	}
}
