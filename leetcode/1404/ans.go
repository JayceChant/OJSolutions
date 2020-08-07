package main

// 0ms
// 2MB
func numSteps(s string) int {
	step := 0
	carry := 0
	i := len(s) - 1
	for i >= 1 {
		count := 0
		for i >= 1 && s[i] == '0' {
			count++
			i--
		}
		step += (carry + 1) * (count - carry)

		count = carry
		for i >= 1 && s[i] == '1' {
			count++
			i--
		}

		if i == 0 {
			if count == 0 {
				break
			}
			count++
		}
		step += 1 + count

		carry = 1
	}
	return step
}
