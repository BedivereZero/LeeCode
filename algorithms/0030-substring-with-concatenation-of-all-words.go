package algorithms

func findSubstring(s string, words []string) []int {
	result := []int{}
	if len(s) == 0 || len(words) == 0 {
		return result
	}
	existence := make([]bool, len(words))
	wordLength := len(words[0])
	for i := 0; i < wordLength; i++ {
		head, curr := i, i
		unmarkAll(existence)
		for head+len(words)*wordLength < len(s)+1 {
			word := s[curr : curr+wordLength]

			full, valid := isFull(existence, words, word), validate(words, word)

			if !valid {
				unmarkAll(existence)
				curr += wordLength
				head = curr
				continue
			}

			if full {
				for isFull(existence, words, word) {
					unmark(existence, words, s[head:head+wordLength])
					head += wordLength
				}
				mark(existence, words, s[curr:curr+wordLength])
				curr += wordLength
				continue
			}

			mark(existence, words, word)
			curr += wordLength

			if head+wordLength*len(words) == curr {
				result = append(result, head)
				unmark(existence, words, s[head:head+wordLength])
				head += wordLength
			}
		}
	}

	return result
}

func mark(existence []bool, words []string, word string) {
	for i, w := range words {
		if w == word && !existence[i] {
			existence[i] = true
			return
		}
	}
}

func unmark(existence []bool, words []string, word string) {
	for i, w := range words {
		if w == word && existence[i] {
			existence[i] = false
			return
		}
	}
}

func unmarkAll(existence []bool) {
	for i := range existence {
		existence[i] = false
	}
}

func isFull(existence []bool, words []string, word string) bool {
	for i, w := range words {
		if w == word && !existence[i] {
			return false
		}
	}
	return true
}

func validate(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}
