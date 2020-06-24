package algorithms

import "fmt"

const (
	intMin = -1 << 31
	intMax = 1<<31 - 1
)

func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > intMin {
			return -dividend
		} else {
			return intMax
		}
	}

	n := 0
	t := divisor
	for t > 0 && t < intMax>>1 || t < 0 && t >= intMin>>1 {
		t = t << 1
		n++
	}
	fmt.Printf("n: %d, t: %d\n", n, t)
	fmt.Printf("dividend: %d, divisor: %d\n", dividend, divisor)

	r := 0
	for n >= 0 && dividend != 0 {
		fmt.Printf("n: %d, dividend: %d, t: %d, r: %d\n", n, dividend, t, r)
		if dividend > 0 && t > 0 && t <= dividend || dividend > 0 && t < 0 && -t <= dividend || dividend < 0 && t > 0 && -t >= dividend || dividend < 0 && t < 0 && t >= dividend {
			// r += 1 << uint(n)
			if dividend > 0 && t > 0 || dividend < 0 && t < 0 {
				r += 1 << uint(n)
				dividend -= t
			} else {
				r -= 1 << uint(n)
				dividend += t
			}
		}
		if dividend > 0 && t > 0 && t > dividend || dividend > 0 && t < 0 && -t > dividend || dividend < 0 && t > 0 && -t < dividend || dividend < 0 && t < 0 && t < dividend {
			t = t >> 1
			n--
		}
	}

	return r
}
