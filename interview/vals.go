package interview

import (
	"testing"
)

var (
	// 左 -> 下 -> 右 -> 上
	// (x,y)
	directions = [4][2]int{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}
)

var (
	RedColor   = "\033[31m"
	ResetColor = "\033[0m"
)

func TestColor(t *testing.T) {
	t.Logf("%s%s%s", RedColor, "hello", ResetColor)
	// fmt.Printf("")
}
