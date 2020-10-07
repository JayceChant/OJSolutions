package main

import (
	"sort"
)

// 104ms
// 6MB
func maxRectangle(words []string) []string {
	lenGroup := map[int]([]string){}
	var lens []int
	for _, word := range words {
		length := len(word)
		lenGroup[length] = append(lenGroup[length], word)
	}
	for length := range lenGroup {
		sort.Strings(lenGroup[length])
		lens = append(lens, length)
	}
	sort.Slice(lens, func(i, j int) bool {
		return lens[i] > lens[j]
	})

	candRows = make([]string, 0)
	candCols = make([]string, 0)
	for _, w := range lens {
		for _, h := range lens {
			rows = lenGroup[w]
			cols = lenGroup[h]
			width = w
			height = h
			candRows = candRows[:0]
			candCols = candCols[:0]
			ans, ok := btRow(0, 0)
			if ok {
				return ans
			}
		}
	}
	return []string{}
}

var (
	rows, cols, candRows, candCols []string
	width, height                  int
)

func btRow(r, c int) (ans []string, ok bool) {
	candRows = append(candRows, "")
ROW_LOOP:
	for _, row := range rows {
		for ci, col := range candCols {
			if row[ci] != col[r] {
				continue ROW_LOOP
			}
		}

		candRows[r] = row
		if r == height-1 && c == width {
			return candRows, true
		}

		if c == width {
			ans, ok = btRow(r+1, c)
		} else {
			ans, ok = btCol(r+1, c)
		}
		if ok {
			return
		}
	}
	candRows = candRows[:r]
	return nil, false
}

func btCol(r, c int) (ans []string, ok bool) {
	candCols = append(candCols, "")
COL_LOOP:
	for _, col := range cols {
		for ri, row := range candRows {
			if col[ri] != row[c] {
				continue COL_LOOP
			}
		}

		candCols[c] = col
		if r == height && c == width-1 {
			return candRows, true
		}
		if r == height {
			ans, ok = btCol(r, c+1)
		} else {
			ans, ok = btRow(r, c+1)
		}
		if ok {
			return
		}
	}
	candCols = candCols[:c]
	return nil, false
}
