package main

import "testing"

var (
	t     = "asfdgdhgjdghfhjtyiopomnfnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnefwghoxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxhjtyiopomnfnnnnnnnnnnnnnnnnnxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxgvsdhijiopomnvzqffhrjtkiuoasfdgdhgjdghpiopomnklkmnogvsdhijiopomnvzqffhrjtkiuoasfdgdhgjdghpiopogfavvrrwtytwutasfdgdhgjdghwrqwsfgghnjkiopkjfcxsyyyyrghjkfhjtyiopomnfnnnnnnnnnnnnnnnnnavvrrwtytwutasfdgdhgjdghwxxxxxxxxxxxxxxxhhdfgfhhdhgsgwrtyurtergggfgdgdfggfgbbnnncvsdfrgrakklwsvomqzgohwrpppjbfvdcvg"
	tests = []string{
		"gdhgjghtyfwghogkiklkngfvvrwutwrqwsfgghnjfcxyyyrghj",
		"fhiogfffdsgfhyfbnnmmmmmerwqwqdyyyyvczcyewefsqaadfg",
	}
)

func Benchmark_isSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSubsequence(tests[0], t)
		isSubsequence(tests[1], t)
	}
}

func Benchmark_isSubsequenceCacheDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSubsequenceDP(tests[0], t)
		isSubsequenceDP(tests[1], t)
	}
}
