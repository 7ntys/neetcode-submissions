func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {return "0"}

	m := len(num1)
	n := len(num2)

	res := make([]int, m+n)

	for i:=m-1;i>=0;i--{
		for j:=n-1;j>=0;j--{
			digit1, digit2 := int(num1[i] - '0'), int(num2[j] - '0')

			prod := digit1 * digit2

			curr := i + j + 1
			carry := i + j

			sum := prod + res[curr]
			res[curr] = sum % 10
			res[carry] += sum / 10
		}
	}

	out := []byte{}
	for i:=0;i<len(res);i++{
		if len(out) == 0 && res[i] == 0 {continue}
		out = append(out, byte(res[i] + '0'))
	}
	return string(out)
}
