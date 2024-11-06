package dynamic

import "testing"

func TestMinEditDistanceWithInsertDeleteReplace(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "eat and sea",
			args: args{
				str1: "eat",
				str2: "sea",
			},
			want: 2,
		},
		{
			name: "horse and hos",
			args: args{
				str1: "horse",
				str2: "hos",
			},
			want: 2,
		},
		{
			name: "lrbb and mqbhcda",
			args: args{
				str1: "lrbb",
				str2: "mqbhcda",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinEditDistanceWithInsertDeleteReplace(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("MinEditDistanceWithInsertDeleteReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}
