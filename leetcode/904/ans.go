package main

// 112ms
// 7.1MB
func totalFruit(tree []int) int {
	max := 0
	var fruit1, fruit2, start int
	lasti1 := -1
	lasti2 := -1

	for i, f := range tree {
		if f == fruit1 {
			lasti1 = i
			continue
		}

		if f == fruit2 {
			lasti2 = i
			continue
		}

		// new fruit, count
		collect := i - start
		if max < collect {
			max = collect
		}

		if lasti1 < lasti2 {
			start = lasti1 + 1
			fruit1 = f
			lasti1 = i
		} else {
			start = lasti2 + 1
			fruit2 = f
			lasti2 = i
		}
	}

	collect := len(tree) - start
	if max < collect {
		max = collect
	}

	return max
}
