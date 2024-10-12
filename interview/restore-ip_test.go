package interview

import (
	"reflect"
	"testing"
)

// 1) 25525511135 -> ["255.255.11.135", "255.255.111.35"]
// 2) 0000 -> ["0.0.0.0"]
// 3) 110110 -> ["1.10.11.0", "1.10.1.10", "1.1.0.110", "11.0.1.10", "11.0.11.0", "110.1.1.0"]
// 4) "010010" -> ["0.10.0.10", "0.100.1.0"]
func TestRestoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "25525511135",
			args: args{s: "25525511135"},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "0000",
			args: args{s: "0000"},
			want: []string{"0.0.0.0"},
		},
		{
			name: "110110",
			args: args{s: "110110"},
			//1.1.0.110 1.10.1.10 1.10.11.0 1.101.1.0 11.0.1.10 11.0.11.0 110.1.1.0
			want: []string{"1.1.0.110", "1.10.1.10", "1.10.11.0", "1.101.1.0", "11.0.1.10", "11.0.11.0", "110.1.1.0"},
		},
		{
			name: "010010",
			args: args{s: "010010"},
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RestoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestoreIpAddressesII(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "25525511135",
			args: args{s: "25525511135"},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "0000",
			args: args{s: "0000"},
			want: []string{"0.0.0.0"},
		},
		{
			name: "110110",
			args: args{s: "110110"},
			//1.1.0.110 1.10.1.10 1.10.11.0 1.101.1.0 11.0.1.10 11.0.11.0 110.1.1.0
			want: []string{"1.1.0.110", "1.10.1.10", "1.10.11.0", "1.101.1.0", "11.0.1.10", "11.0.11.0", "110.1.1.0"},
		},
		{
			name: "010010",
			args: args{s: "010010"},
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			name: "101@023",
			args: args{s: "101@023"},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RestoreIpAddressesII(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestoreIpAddressesII() = %v, want %v", got, tt.want)
			}
		})
	}
}
