package main

import "strings"

func convert(s string, numRows int) string {
	n := len(s)
	if n <= numRows || numRows == 1 {
		return s
	}
	base := numRows - 1

	numCols := make([]int, base)
	numCols[0] = (n + base<<1 - 1) / (base << 1)
	midCols := n / base
	rest := n % base
	if midCols%2 == 0 {
		for i := 1; i < base; i++ {
			numCols[i] = midCols
			if rest > i {
				numCols[i]++
			}
		}
	} else {
		for i := base - 1; i >= 1; i-- {
			numCols[i] = midCols
			if rest > base-i {
				numCols[i]++
			}
		}
	}
	//fmt.Println(numCols)
	cols := make([]int, numRows)

	ans := make([]byte, n)
	row := 0
	down := false
	to := 0
	for from := range s {
		//fmt.Println(to)
		ans[to] = s[from]
		if row == 0 || row == base {
			down = !down
		}
		if down {
			to += (numCols[row] - cols[row] + cols[row+1])
			cols[row]++
			row++
		} else {
			to -= (cols[row] + numCols[row-1] - cols[row-1])
			cols[row]++
			row--
		}
	}
	return string(ans)
}

func convert2(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	byteSlices := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		byteSlices[i] = make([]byte, 0)
	}
	row := 0
	direction := 1
	for i := range s {
		byteSlices[row] = append(byteSlices[row], s[i])
		row += direction
		if row == 0 || row == numRows-1 {
			direction = -direction
		}
	}
	sb := strings.Builder{}
	for i := 0; i < numRows; i++ {
		sb.Write(byteSlices[i])
	}
	return sb.String()
}
