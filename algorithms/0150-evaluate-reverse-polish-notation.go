package algorithms

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	depth := 0
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			stack[depth-2] += stack[depth-1]
			depth--
		case "-":
			stack[depth-2] -= stack[depth-1]
			depth--
		case "*":
			stack[depth-2] *= stack[depth-1]
			depth--
		case "/":
			stack[depth-2] /= stack[depth-1]
			depth--
		default:
			n, _ := strconv.Atoi(tokens[i])
			if len(stack) > depth {
				stack[depth] = n
			} else {
				stack = append(stack, n)
			}
			depth++
		}
	}
	return stack[0]
}
