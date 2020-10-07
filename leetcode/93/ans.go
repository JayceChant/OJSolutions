package main

// 0ms
// 2.1MB
func restoreIpAddresses(s string) []string {
	str = s
	length = len(s)
	cand = make([]int, 3)
	ans = make([]string, 0)
	backtracking(0, 0)
	return ans
}

const (
	full = "255"
)

var (
	str    string
	length int
	cand   []int
	ans    []string
)

// [st, ed)
func backtracking(i int, st int) {
	//fmt.Println("i/st:", i, st)
	if length-st > 3*(4-i) ||
		length-st < (4-i) {
		return
	}

	if i == 3 {
		if length-st > 1 && str[st] == '0' {
			// no leading 0
			return
		}
		if length-st == 3 {
			for l := 0; l < 3; l++ {
				cmp := int8(str[st+l] - full[l])
				if cmp > 0 {
					return
				}

				if cmp < 0 {
					break
				}

				// cmp == 0
				// continue (compare next byte)
			}
		}

		a, b, c := cand[0], cand[1], cand[2]
		ans = append(ans, str[:a]+"."+str[a:b]+"."+str[b:c]+"."+str[c:])
		//fmt.Println("ans!")
		return
	}

	var cmp int8 = 1
	if length-st >= 3 {
		cmp = 0
	}
	for j := 0; j < 2; j++ {
		ed := st + j + 1
		cand[i] = ed
		backtracking(i+1, ed)

		if str[st] == '0' {
			return
		}

		if cmp == 0 {
			cmp = int8(str[st+j] - full[j])
		}
	}

	if cmp < 0 || (cmp == 0 && str[st+2] <= full[2]) {
		cand[i] = st + 3
		backtracking(i+1, st+3)
	}
}
