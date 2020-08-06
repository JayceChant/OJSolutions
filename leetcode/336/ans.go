package main

// 48ms 100%
// 6.7MB 100%
func palindromePairs(words []string) [][]int {
	mirror := map[string]int{}
	emptyJ := -1
	for i, word := range words {
		if word == "" {
			emptyJ = i
			continue
		}
		mirror[reverse(word)] = i
	}

	ans := make([][]int, 0)
	for i, word := range words {
		if i == emptyJ {
			continue
		}

		if j, ok := mirror[word]; ok && i != j {
			ans = append(ans, []int{i, j})
		}

		wlen := len(word)
		for piv := 1; piv < wlen; piv++ {
			half := piv / 2

			k := 0
			for ; k < half; k++ {
				if word[k] != word[piv-1-k] {
					break
				}
			}
			if k == half {
				if j, ok := mirror[word[piv:]]; ok {
					ans = append(ans, []int{j, i})
				}
			}

			k = 0
			offset := wlen - piv
			for ; k < half; k++ {
				if word[offset+k] != word[wlen-1-k] {
					break
				}
			}
			if k == half {
				if j, ok := mirror[word[:offset]]; ok {
					ans = append(ans, []int{i, j})
				}
			}
		}

		if emptyJ >= 0 {
			half := wlen / 2

			k := 0
			for ; k < half; k++ {
				if word[k] != word[wlen-1-k] {
					break
				}
			}
			if k == half {
				ans = append(ans, []int{i, emptyJ}, []int{emptyJ, i})
			}
		}
	}
	return ans
}

// benchmark with str len 312: 2992347	       388 ns/op	     640 B/op	       2 allocs/op
func reverse(word string) string {
	buf := []byte(word)
	maxIdx := len(word) - 1
	half := len(word) / 2
	for i := 0; i < half; i++ {
		buf[i], buf[maxIdx-i] = buf[maxIdx-i], buf[i]
	}
	return string(buf)
}

// TODO: try manacher

// benchmark with str len 312: 1655077	       723 ns/op	     320 B/op	       1 allocs/op
// func reverse2(word string) string {
// 	sb := strings.Builder{}
// 	sb.Grow(len(word))
// 	for i := len(word) - 1; i >= 0; i-- {
// 		sb.WriteByte(word[i])
// 	}
// 	return sb.String()
// }
