package main

func strStr(haystack string, needle string) int {
	nSize := len(needle)
	hSize := len(haystack)
STLOOP:
	for st := 0; st <= hSize-nSize; st++ {
		for i := 0; i < nSize; i++ {
			if haystack[st+i] != needle[i] {
				continue STLOOP
			}
		}
		return st
	}
	return -1
}
