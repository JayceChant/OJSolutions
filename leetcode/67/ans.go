package main

// 0ms, 100%
// 2.2MB
func addBinary(a string, b string) string {
	ia := len(a) - 1
	ib := len(b) - 1
	if ia < ib {
		a, ia, b, ib = b, ib, a, ia
	}

	buf := make([]byte, ia+2)
	var carry byte = 0
	var digit byte
	for ib >= 0 {
		digit = a[ia] - '0' + b[ib] - '0' + carry
		buf[ia+1] = '0' + digit&1
		carry = digit >> 1
		ia--
		ib--
	}

	for ia >= 0 {
		digit = a[ia] - '0' + carry
		buf[ia+1] = '0' + digit&1
		carry = digit >> 1
		ia--
	}
	buf[0] = '0' + carry

	if buf[0] == '0' {
		return string(buf[1:])
	}
	return string(buf)
}
