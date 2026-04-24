func checkValidString(s string) bool {
    leftMin, leftMax := 0, 0

	for i:=0;i<len(s);i++{
		if s[i] == '(' {
			leftMin++
			leftMax++
		} else if s[i] == ')' {
			leftMin--
			leftMax--
		} else {
			leftMin--
			leftMax++
		}

		if leftMin < 0 {leftMin = 0}
		if leftMax < 0 {return false}
	}
	return leftMin == 0
}
