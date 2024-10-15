package str

import "testing"

func TestIsLongPressName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"alex", "aaleex"}, true},
		{"2", args{"saeed", "ssaaedd"}, false},
		{"3", args{"leelee", "lleeelee"}, true},
		{"4", args{"laiden", "laiden"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLongPressName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("IsLongPressName() = %v, want %v", got, tt.want)
			}
		})
	}
}
