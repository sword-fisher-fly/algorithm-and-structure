package math

// ref: https://leetcode.com/problems/broken-calculator/discuss/234484/JavaC%2B%2BPython-Change-Y-to-X-in-1-Line
//  On a broken calculator that has a number showing on its display, we can perform two operations:

// - Double: Multiply the number on the display by 2, or;
// - Decrement: Subtract 1 from the number on the display.

// Initially, the calculator is displaying the number X.

// Return the minimum number of operations needed to display the number Y.

// Example 1:

// Input: X = 2, Y = 3
// Output: 2
// Explanation: Use double operation and then decrement operation {2 -> 4 -> 3}.
//

func BrokenCalculator(x, y int) int {
	res := 0
	for x < y {
		res += y%2 + 1
		y = (y + 1) / 2
	}
	return res + x - y
}
