package main

import "fmt"

// ClimbStairs counts the distinct ways to climb n steps taking 1 or 2 at a
// time. ways(n) = ways(n-1) + ways(n-2) — Fibonacci — carried in two
// variables instead of an array.
func ClimbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	prev, cur := 1, 1 // ways(0), ways(1)
	for i := 2; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}

func main() {
	fmt.Println(ClimbStairs(5))
}
