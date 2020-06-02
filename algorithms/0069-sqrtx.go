package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	l, r := 0, x
	for l < r - 1 {
		m := (l + r) /2
		if m * m == x {
			return m
		}
		if m * m < x {
			l = m
		} else {
			r = m
		}
	}
	return l
}


func main() {
	fmt.Println(mySqrt(8))
}
