package main

// 4ms
// 3.3MB
func restoreString(s string, indices []int) string {
	length := len(s)
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[indices[i]] = s[i]
	}
	return string(buf)
}
