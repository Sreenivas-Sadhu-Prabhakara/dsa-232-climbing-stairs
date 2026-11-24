package main

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := []struct{ n, want int }{
		{0, 1}, // one way: stand still
		{1, 1},
		{2, 2},
		{3, 3},
		{5, 8},
		{10, 89},
		{45, 1836311903}, // the LC upper bound still fits an int
	}
	for _, c := range cases {
		if got := ClimbStairs(c.n); got != c.want {
			t.Errorf("ClimbStairs(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}
