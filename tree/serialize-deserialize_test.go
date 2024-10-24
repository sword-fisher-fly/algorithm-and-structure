package tree

import (
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	//    1
	//   / \
	//   2 6
	//  /\
	// 3 4
	// /
	// 5
	tests := []struct {
		name string
		c    Codec
		args args
		want string
	}{
		{
			name: "case 1",
			c:    Codec{},
			args: args{root: level4RandomTree.Root()},
			want: "1,2,3,5,null,null,null,4,null,null,6,null,null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Codec{}
			if got := c.serialize(tt.args.root); got != tt.want {
				t.Errorf("Codec.serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		c    Codec
		args args
		want string
	}{
		{
			name: "case 1",
			c:    Codec{},
			args: args{data: "1,2,3,5,null,null,null,4,null,null,6,null,null"},
			want: "1,2,3,5,null,null,null,4,null,null,6,null,null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Codec{}
			if got := c.deserialize(tt.args.data); c.serialize(got) != tt.want {
				t.Errorf("Codec.deserialize() = %v, want %v", c.serialize(got), tt.want)
			}
		})
	}
}
