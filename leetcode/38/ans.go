package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 0ms, 100%
// 2.1MB, 85.71%
func countAndSaySB(n int) string {
	to := strings.Builder{}
	to.WriteByte('1')
	for i := 2; i <= n; i++ {
		from := to.String()
		to.Reset()
		count := 1
		for j := 1; j < len(from); j++ {
			if from[j] == from[j-1] {
				count++
				continue
			}

			to.WriteString(strconv.Itoa(count))
			to.WriteByte(from[j-1])
			count = 1
		}
		to.WriteString(strconv.Itoa(count))
		to.WriteByte(from[len(from)-1])
	}
	return to.String()
}

// 0ms, 100%
// 2.0MB, 100%
func countAndSay(n int) string {
	from := make([]byte, 0)
	to := []byte{'1'}
	for i := 2; i <= n; i++ {
		from, to = to, from[:0]
		var count byte = '1'
		for j := 1; j < len(from); j++ {
			if from[j] == from[j-1] {
				count++
				continue
			}

			to = append(to, count)
			to = append(to, from[j-1])
			count = '1'
		}
		to = append(to, count)
		to = append(to, from[len(from)-1])
	}
	return string(to)
}

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(countAndSaySB(i))
	}
}
