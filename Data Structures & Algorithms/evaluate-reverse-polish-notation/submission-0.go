func evalRPN(tokens []string) int {
	stack := []int{}

	for i := range tokens {
		token := tokens[i]

		a, err := strconv.Atoi(token)

		if err != nil {
			if len(stack) < 2 {return -1}
			n := len(stack)
			a, b := stack[n-2], stack[n-1]
			stack = stack[:n-2]
			// Is an operator : 
			switch token {
				case "*":
					stack = append(stack, a * b)
				case "-":
					stack = append(stack, a - b)
				case "+":
					stack = append(stack, a + b)
				case "/":
					if b == 0 {return -1}
					stack = append(stack, a / b)
			}
		} else {
			stack = append(stack, a)
		}
	}
	return stack[0]
}
