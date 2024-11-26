package huawei

import (
	"reflect"
	"testing"
)

func TestParseArguments(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "without quote",
			args: args{"xcopy /s c:// d://e"},
			want: []string{"xcopy", "/s", "c://", "d://e"},
		},
		{
			name: "with quote",
			args: args{"l \"c\" d"},
			want: []string{"l", "c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseArguments(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArguments() = %v, want %v", got, tt.want)
			}
		})
	}
}
