package main

// 0ms
// 4.9MB
func minFlips(target string) int {
	length := len(target)
	count := 0
	i := 0

	for i < length && target[i] == '0' {
		i++
	}

	for i < length {
		for i < length && target[i] == '1' {
			i++
		}
		count++

		if i == length {
			break
		}

		for i < length && target[i] == '0' {
			i++
		}
		count++
	}
	return count
}
