package str

import "testing"

func TestCleanPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "with  both dot and dot-dot path part",
			args: args{path: "/a/./b/../../c/"},
			want: "/c",
		},
		{
			name: "with dot path part",
			args: args{path: "/a/b/./c/"},
			want: "/a/b/c",
		},
		{
			name: "with dot-dot path part",
			args: args{path: "/a/b/../c/d/../e/f//"},
			want: "/a/c/e/f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanPath(tt.args.path); got != tt.want {
				t.Errorf("CleanPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
