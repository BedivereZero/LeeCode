package algorithms

import "bytes"

func multiply(num1 string, num2 string) string {
	resultArray := make([]int, len(num1)+len(num2))
	for i := 0; i < len(num2); i++ {
		add(resultArray, multiplySingle(num1, num2[len(num2)-1-i]), i)
	}
	var buff bytes.Buffer
	leadingZero := true
	for i := 0; i < len(resultArray); i++ {
		if resultArray[len(resultArray)-1-i] == 0 && leadingZero {
			continue
		}
		if leadingZero && resultArray[len(resultArray)-1-i] != 0 {
			leadingZero = false
		}
		buff.WriteByte(byte(resultArray[len(resultArray)-1-i] + '0'))
	}
	if buff.Len() == 0 {
		buff.WriteByte('0')
	}
	return buff.String()
}

func multiplySingle(a string, b byte) []int {
	bb := int(b - '0')
	c := make([]int, len(a)+1)
	carry := 0
	for i := 0; i < len(a); i++ {
		c[i] = int(a[len(a)-i-1]-'0')*bb + carry
		carry, c[i] = c[i]/10, c[i]%10
	}
	c[len(c)-1] = carry
	return c
}

func add(a []int, b []int, c int) {
	carry := 0
	for i := 0; i < len(b); i++ {
		a[c+i] = a[c+i] + b[i] + carry
		a[c+i], carry = a[c+i]%10, a[c+i]/10
	}
	for i := len(b) + c; i < len(a); i++ {
		a[i] = a[i] + carry
		a[i], carry = a[i]%10, a[i]/10
	}
}
