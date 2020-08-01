package algorithms

import (
	"bytes"
)

func reverseWords(s string) string {
	var buff bytes.Buffer
	for h, t := len(s), len(s); t > 0; t = h {
		// find word end
		for ; t > 0 && s[t-1] == ' '; t-- {
		}
		// find word begin
		for h = t; h > 0 && s[h-1] != ' '; h-- {
		}
		// store word, if exists
		if h < t {
			if buff.Len() > 0 {
				buff.WriteByte(' ')
			}
			buff.WriteString(s[h:t])
		}
	}
	return buff.String()
}
