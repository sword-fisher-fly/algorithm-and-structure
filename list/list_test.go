package list

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	cs := []struct {
		name     string
		nodeVals []int64
		want     []int64
	}{
		{
			name:     "test1",
			nodeVals: []int64{1, 2, 3, 4, 5},
			want:     []int64{1, 2, 3, 4, 5},
		},
	}
	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			l := NewList()
			for _, v := range c.nodeVals {
				l.Add(v)
			}

			if got := l.ToArray(); !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func TestListAddNode(t *testing.T) {
	l := NewList()

	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)

	l.AddNode(node1)
	l.AddNode(node2)
	l.AddNode(node3)
	l.AddNode(node4)
	l.AddNode(node5)

	if got := l.ToArray(); !reflect.DeepEqual(got, []int64{1, 2, 3, 4, 5}) {
		t.Errorf("got %v, want %v", got, []int64{1, 2, 3, 4, 5})
	}

	l.RemoveNode(node3)

	if got := l.ToArray(); !reflect.DeepEqual(got, []int64{1, 2, 4, 5}) {
		t.Errorf("got %v, want %v", got, []int64{1, 2, 4, 5})
	}

	if l.Tail.Value != 5 {
		t.Errorf("got %v, want %v", l.Tail.Value, 5)
	}

	l.RemoveNode(node5)
	if l.Tail.Value != 4 {
		t.Errorf("got %v, want %v", l.Tail.Value, 4)
	}
}

func TestListRemove(t *testing.T) {
	testcases := []struct {
		name   string
		arr    []int64
		remove int64
		want   []int64
	}{
		{
			name:   "remove only one",
			remove: 4,
			arr:    []int64{1, 2, 3, 4, 5},
			want:   []int64{1, 2, 3, 5},
		},
		{
			name:   "remove head",
			arr:    []int64{1, 2, 3, 4, 5},
			remove: 1,
			want:   []int64{2, 3, 4, 5},
		},
		{
			name:   "remove tail",
			arr:    []int64{1, 2, 3, 4, 5},
			remove: 5,
			want:   []int64{1, 2, 3, 4},
		},
		{
			name:   "remove two",
			arr:    []int64{1, 2, 3, 3, 4, 5},
			remove: 3,
			want:   []int64{1, 2, 4, 5},
		},
	}

	for _, c := range testcases {
		t.Run(c.name, func(t *testing.T) {
			l := NewList()
			l.InitializeFromArray(c.arr)
			l.Remove(c.remove)

			if got := l.ToArray(); !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})

	}
}

func TestList_Reverse(t *testing.T) {
	tests := []struct {
		name string
		arr  []int64
		want []int64

		headVal int64
		tailVal int64
	}{
		{
			name: "one item",
			arr:  []int64{1},
			want: []int64{1},

			headVal: 1,
			tailVal: 1,
		},
		{
			name: "two items",
			arr:  []int64{1, 2},
			want: []int64{2, 1},

			headVal: 2,
			tailVal: 1,
		},
		{
			name: "10 items",
			arr:  []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},

			headVal: 10,
			tailVal: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList()
			l.InitializeFromArray(tt.arr)

			l.Reverse()
			if got := l.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
