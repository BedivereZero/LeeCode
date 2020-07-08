package main

import (
	"bytes"
	"fmt"
)

func main() {
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth := 16
	for _, line := range fullJustify(words, maxWidth) {
		fmt.Println(line, "#")
	}
}

func fullJustify(words []string, maxWidth int) []string {
	solution := []string{}
	for head := 0; head < len(words); {
		tail, width := head, 0
		for tail < len(words) {
			if tail > head {
				width++
			}
			if width+len(words[tail]) > maxWidth {
				break
			}
			width += len(words[tail])
			tail++
		}

		// fmt.Println("Line words:", words[head:tail])

		var line bytes.Buffer
		if tail < len(words) {
			line.WriteString(words[head])
			spaceCount := tail - head - 1
			spaceTotal := genSpaceTotalCount(words, head, tail, maxWidth)
			// fmt.Println("Space:", spaceCount, ", Space total:", spaceTotal)
			for pos, word := range words[head+1 : tail] {
				for i := 0; i < spaceTotal/spaceCount; i++ {
					line.WriteByte(' ')
				}
				if pos < spaceTotal%spaceCount {
					line.WriteByte(' ')
				}
				line.WriteString(word)
			}
			for line.Len() < maxWidth {
				line.WriteByte(' ')
			}
		} else {
			line.WriteString(words[head])
			for _, word := range words[head+1 : tail] {
				line.WriteByte(' ')
				line.WriteString(word)
			}
			for line.Len() < maxWidth {
				line.WriteByte(' ')
			}
		}

		// fmt.Println("Line's length:", line.Len())
		solution = append(solution, line.String())

		head = tail
	}

	return solution
}

func genSpaceTotalCount(words []string, head int, tail int, maxWidth int) int {
	for _, word := range words[head:tail] {
		maxWidth -= len(word)
	}
	return maxWidth
}
