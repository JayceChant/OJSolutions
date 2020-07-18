package main

// 0ms, 100%
// 2MB, 100%
func isInterleave(s1 string, s2 string, s3 string) bool {
	size1 := len(s1)
	size2 := len(s2)
	if size1+size2 != len(s3) {
		return false
	}

	if size1 > size2 {
		// make s1 the shorter one
		s1, size1, s2, size2 = s2, size2, s1, size1
	}

	dpLastRow := make([]bool, size1+1)
	dpCurrRow := make([]bool, size1+1)
	dpLastRow[size1] = true

	increase := true // flag for early pruning, true for 1st row
	for i2 := size2; i2 >= 0; i2-- {
		dpLastRow, dpCurrRow = dpCurrRow, dpLastRow
		for i1 := size1; i1 >= 0; i1-- {
			// i1+1 interleave i2 match, and i1 match
			if i1 < size1 && dpCurrRow[i1+1] && s1[i1] == s3[i1+i2] {
				dpCurrRow[i1] = true
				increase = true
				continue
			}

			// i1 interleave i2+1 match, and i2 match
			if i2 < size2 && dpLastRow[i1] && s2[i2] == s3[i1+i2] {
				dpCurrRow[i1] = true
				increase = true
				continue
			}

			if i1 < size1 || i2 < size2 {
				dpCurrRow[i1] = false
			}
		}
		//fmt.Println(dpCurrRow)
		if !increase {
			return false
		}
		increase = false
	}
	return dpCurrRow[0]
}
