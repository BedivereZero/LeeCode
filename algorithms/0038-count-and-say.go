package algorithms

import (
	"bytes"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		last := countAndSay(n - 1)
		first := last[0]
		num := 1
		var buffer bytes.Buffer
		for i := 1; i < len(last); i++ {
			if last[i] == first {
				num++
			} else {
				buffer.WriteString(strconv.Itoa(num))
				buffer.WriteString(string(first))
				first = last[i]
				num = 1
			}
		}
		buffer.WriteString(strconv.Itoa(num))
		buffer.WriteString(string(first))
		return buffer.String()
	}
}
