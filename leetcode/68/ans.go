package main

import "strings"

// 0ms, 100%
// 2.1MB, 100%
func fullJustify(words []string, maxWidth int) []string {
	length := len(words)
	rows := make([]string, 0)
	spaces := strings.Repeat(" ", maxWidth-1)
	sb := strings.Builder{}

	end := 0
	for end < length {
		start := end
		// word length would not exceed the max width
		sb.WriteString(words[start])
		width := maxWidth - len(words[end])
		end++

		for end < length && width > len(words[end]) {
			width -= (len(words[end]) + 1)
			end++
		}

		if end-start == 1 || end == length {
			// left-aligned
			for i := start + 1; i < end; i++ {
				sb.WriteByte(' ')
				sb.WriteString(words[i])
			}
			sb.WriteString(spaces[:width])
		} else {
			spCount := end - start - 1
			spLen := width/spCount + 1
			spMore := start + width%spCount
			for i := start + 1; i < end; i++ {
				sb.WriteString(spaces[:spLen])
				if i <= spMore {
					sb.WriteByte(' ')
				}
				sb.WriteString(words[i])
			}
		}

		rows = append(rows, sb.String())
		sb.Reset()
	}

	return rows
}
