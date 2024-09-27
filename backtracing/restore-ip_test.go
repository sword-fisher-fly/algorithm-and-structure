package backtracking

import (
	"reflect"
	"testing"
)

func TestRestoreIPAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1111",
			args: args{s: "1111"},
			want: []string{"1.1.1.1"},
		},
		{
			name: "25525511135",
			args: args{s: "25525511135"},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "0102030",
			args: args{s: "0102030"},
			want: []string{"0.10.20.30", "0.10.203.0", "0.102.0.30"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RestoreIPAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestoreIPAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
