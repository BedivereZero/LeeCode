package algorithms

import "fmt"

func fractionToDecimal(numerator int, denominator int) string {
	nagtive := numerator*denominator < 0

	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}

	integer := numerator / denominator
	remainder := numerator % denominator

	if remainder == 0 {
		if nagtive {
			return fmt.Sprintf("-%d", integer)
		}
		return fmt.Sprintf("%d", integer)
	}

	decimalBytes := []byte{}
	remainders := []int{}
	for remainder != 0 && !exists(remainders, remainder) {
		decimalBytes = append(decimalBytes, byte(remainder*10/denominator+'0'))
		remainders = append(remainders, remainder)
		remainder = remainder * 10 % denominator
	}

	if remainder == 0 {
		if nagtive {
			return fmt.Sprintf("-%d.%s", integer, decimalBytes)
		}
		return fmt.Sprintf("%d.%s", integer, decimalBytes)
	}

	repeatHead := len(decimalBytes) - 1
	for ; repeatHead >= 0 && remainders[repeatHead] != remainder; repeatHead-- {
	}

	if repeatHead > 0 {
		if nagtive {
			return fmt.Sprintf("-%d.%s(%s)", integer, decimalBytes[:repeatHead], decimalBytes[repeatHead:])
		}
		return fmt.Sprintf("%d.%s(%s)", integer, decimalBytes[:repeatHead], decimalBytes[repeatHead:])
	}
	if nagtive {
		return fmt.Sprintf("-%d.(%s)", integer, decimalBytes)
	}
	return fmt.Sprintf("%d.(%s)", integer, decimalBytes)
}

func exists(nums []int, n int) bool {
	for _, v := range nums {
		if v == n {
			return true
		}
	}
	return false
}
