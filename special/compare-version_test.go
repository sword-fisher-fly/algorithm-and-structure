package special

import (
	"testing"
)

func TestCompareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal case 1",
			args: args{
				version1: "1.1.1",
				version2: "1.1.1",
			},
			want: 0,
		},
		{
			name: "equal case 2",
			args: args{
				version1: "1.1.0",
				version2: "1.1",
			},
			want: 0,
		},
		{
			name: "equal case 3",
			args: args{
				version1: "1.1",
				version2: "1.1.0.0",
			},
			want: 0,
		},
		{
			name: "greater case 4",
			args: args{
				version1: "1.1.1",
				version2: "1.1.0",
			},
			want: 1,
		},
		{
			name: "less case 5",
			args: args{
				version1: "1.1.0",
				version2: "1.1.1",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("CompareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareVersionII(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal case 1",
			args: args{
				version1: "1.1.1",
				version2: "1.1.1",
			},
			want: 0,
		},
		{
			name: "equal case 2",
			args: args{
				version1: "1.1.0",
				version2: "1.1",
			},
			want: 0,
		},
		{
			name: "equal case 3",
			args: args{
				version1: "1.1",
				version2: "1.1.0.0",
			},
			want: 0,
		},
		{
			name: "greater case 4",
			args: args{
				version1: "1.1.1",
				version2: "1.1.0",
			},
			want: 1,
		},
		{
			name: "less case 5",
			args: args{
				version1: "1.1.0",
				version2: "1.1.1",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareVersionII(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("CompareVersionII() = %v, want %v", got, tt.want)
			}
		})
	}
}
