package str

import "testing"

func TestAddBigNum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s1: "11",
				s2: "1234",
			},
			want: "1245",
		},
		{
			name: "case 2",
			args: args{
				s1: "2345",
				s2: "56",
			},
			want: "2401",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBigNum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("AddBigNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
