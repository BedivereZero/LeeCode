package algorithms

import "bytes"

func addStrings(num1 string, num2 string) string {
	// result
	var buff []int
	// carry from previous stage
	var carry = 0
	if len(num1) < len(num2) {
		buff = make([]int, len(num2))
	} else {
		buff = make([]int, len(num1))
	}
	// offset from tail
	var offset = 1
	for offset <= len(num1) && offset <= len(num2) {
		sum := int(num1[len(num1)-offset]-'0') + int(num2[len(num2)-offset]-'0') + carry
		buff[len(buff)-offset] = sum % 10
		carry = sum / 10
		offset++
	}
	for offset <= len(num1) {
		sum := int(num1[len(num1)-offset]-'0') + carry
		buff[len(buff)-offset] = sum % 10
		carry = sum / 10
		offset++
	}
	for offset <= len(num2) {
		sum := int(num2[len(num2)-offset]-'0') + carry
		buff[len(buff)-offset] = sum % 10
		carry = sum / 10
		offset++
	}
	var result bytes.Buffer
	if carry == 1 {
		result.WriteByte('1')
	}
	for _, each := range buff {
		result.WriteByte(byte(each + '0'))
	}
	return result.String()
}
