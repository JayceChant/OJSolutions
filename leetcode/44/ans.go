package main

// '?' 可以匹配任何单个字符。
// '*' 可以匹配任意字符串（包括空字符串）。

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	lastDp := make([]bool, len(p)+1)
	curDp := make([]bool, len(p)+1)
	for is := len(s); is >= 0; is-- {
		for ip := len(p) - 1; ip >= 0; ip-- {
			if p[ip] == '*' {
				curDp[ip] = (ip+1 < len(p) && curDp[ip+1]) ||
					(is < len(s) && lastDp[ip]) || (is == len(s) && ip+1 == len(p))
			} else {
				curDp[ip] = (is < len(s) && (p[ip] == '?' || s[is] == p[ip])) &&
					(lastDp[ip+1] || (is+1 == len(s) && ip+1 == len(p)))
			}
		}
		lastDp, curDp = curDp, lastDp
	}
	//fmt.Println(lastDp)
	return lastDp[0]
}
