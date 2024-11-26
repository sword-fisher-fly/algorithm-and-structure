package special

import (
	"testing"
)

func TestDecodeNumber(t *testing.T) {
	type args struct {
		nums string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1 simple",
			args: args{
				nums: "12",
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: "31717126241541717",
			},
			want: 192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeNumber(tt.args.nums); got != tt.want {
				t.Errorf("DecodeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeNumberII(t *testing.T) {
	type args struct {
		nums string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1 simple",
			args: args{
				nums: "12",
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: "31717126241541717",
			},
			want: 192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeNumberII(tt.args.nums); got != tt.want {
				t.Errorf("DecodeNumberII() = %v, want %v", got, tt.want)
			}
		})
	}
}
