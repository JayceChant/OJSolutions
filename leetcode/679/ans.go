package main

import (
	"math/big"
)

const (
	fullMask = 1<<4 - 1
)

// 4ms (min 0ms)
// 3040kb (min 1940)
func judgePoint24Rat(nums []int) bool {
	var dp2 = map[int](map[rat]bool){}
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			a := &rat{n: nums[i], d: 1}
			b := &rat{n: nums[j], d: 1}
			mask := 1<<i | 1<<j
			result := make(map[rat]bool)
			// +
			result[*(a.add(b))] = true
			// -
			result[*(a.sub(b))] = true
			result[*(b.sub(a))] = true
			// *
			result[*(a.mul(b))] = true
			// /
			result[*(a.div(b))] = true
			result[*(b.div(a))] = true

			dp2[mask] = result
		}
	}
	//fmt.Println(dp2)

	for mask1, res1 := range dp2 {
		if mask1&1 == 0 {
			// make sure mask1 is odd, mask2 is even to de-duplicate
			continue
		}
		res2 := dp2[mask1^fullMask]
		for a := range res1 {
			for b := range res2 {
				// + - *
				if c := a.add(&b); c.isInt() && c.n == 24 {
					return true
				}
				if c := a.sub(&b); c.isInt() && c.n == 24 {
					return true
				}
				if c := b.sub(&a); c.isInt() && c.n == 24 {
					return true
				}
				if c := a.mul(&b); c.isInt() && c.n == 24 {
					return true
				}
				// /
				if a.n != 0 && b.n != 0 {
					if c := a.div(&b); c.isInt() && c.n == 24 {
						return true
					}
					if c := b.div(&a); c.isInt() && c.n == 24 {
						return true
					}
				}
			}
		}
	}

	var dp3 = map[int](map[rat]bool){}
	for i := 0; i < 4; i++ {
		mask1 := 1 << i
		a := &rat{n: nums[i], d: 1}
		for mask2, res := range dp2 {
			if mask1&mask2 != 0 {
				continue
			}
			mask := mask1 | mask2
			result := dp3[mask]
			if result == nil {
				result = make(map[rat]bool)
				dp3[mask] = result
			}
			for b := range res {
				// + - *
				result[*(b.add(a))] = true
				result[*(a.sub(&b))] = true
				result[*(b.sub(a))] = true
				result[*(b.mul(a))] = true
				result[*(b.div(a))] = true
				if b.n != 0 {
					result[*(a.div(&b))] = true
				}
			}
		}
	}
	//fmt.Println(dp3)

	for i := 0; i < 4; i++ {
		mask1 := 1 << i
		a := &rat{n: nums[i], d: 1}
		res := dp3[mask1^fullMask]
		for b := range res {
			// + - *
			if c := b.add(a); c.isInt() && c.n == 24 {
				return true
			}
			if c := a.sub(&b); c.isInt() && c.n == 24 {
				return true
			}
			if c := b.sub(a); c.isInt() && c.n == 24 {
				return true
			}
			if c := b.mul(a); c.isInt() && c.n == 24 {
				return true
			}
			// /
			if c := b.div(a); c.isInt() && c.n == 24 {
				return true
			}
			if b.n != 0 {
				if c := a.div(&b); c.isInt() && c.n == 24 {
					return true
				}
			}
		}
	}
	return false
}

type rat struct {
	n, d int
}

func (a *rat) add(b *rat) *rat {
	c := &rat{
		n: a.n*b.d + b.n*a.d,
		d: a.d * b.d,
	}
	c.normal()
	return c
}

func (a *rat) sub(b *rat) *rat {
	c := &rat{
		n: a.n*b.d - b.n*a.d,
		d: a.d * b.d,
	}
	c.normal()
	return c
}

func (a *rat) mul(b *rat) *rat {
	c := &rat{
		n: a.n * b.n,
		d: a.d * b.d,
	}
	c.normal()
	return c
}

func (a *rat) div(b *rat) *rat {
	c := &rat{
		n: a.n * b.d,
		d: a.d * b.n,
	}
	c.normal()
	return c
}

func (a *rat) normal() {
	if a.n == 0 {
		a.d = 1
		return
	}
	f := gcd(a.n, a.d)
	if f > 1 {
		a.n /= f
		a.d /= f
	}
}

func (a *rat) isInt() bool {
	return a.d == 1
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

const (
	// normal min result is 1/9/9/9 = 0.00137...
	precision = 1e-5
)

// 0ms
// 2196kB (min 1904)
func judgePoint24(nums []int) bool {
	var dp2 = map[int](map[float32]bool){}
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			a := float32(nums[i])
			b := float32(nums[j])
			mask := 1<<i | 1<<j
			result := make(map[float32]bool)
			// +
			result[a+b] = true
			// -
			result[a-b] = true
			result[b-a] = true
			// *
			result[a*b] = true
			// /
			result[a/b] = true
			result[b/a] = true

			dp2[mask] = result
		}
	}
	// fmt.Println(dp2)

	for mask1, res1 := range dp2 {
		if mask1&1 == 0 {
			// make sure mask1 is odd, mask2 is even to de-duplicate
			continue
		}
		res2 := dp2[mask1^fullMask]
		for a := range res1 {
			for b := range res2 {
				// + - *
				if near24(a+b) || near24(a-b) || near24(b-a) || near24(a*b) {
					return true
				}
				// /
				if a != 0 && b != 0 && (near24(a/b) || near24(b/a)) {
					return true
				}
			}
		}
	}

	var dp3 = map[int](map[float32]bool){}
	for i := 0; i < 4; i++ {
		mask1 := 1 << i
		a := float32(nums[i])
		for mask2, res := range dp2 {
			if mask1&mask2 != 0 {
				continue
			}
			mask := mask1 | mask2
			result := dp3[mask]
			if result == nil {
				result = make(map[float32]bool)
				dp3[mask] = result
			}
			for b := range res {
				// + - *
				result[a+b] = true
				result[a-b] = true
				result[b-a] = true
				result[a*b] = true
				result[b/a] = true
				if b != 0 {
					result[a/b] = true
				}
			}
		}
	}
	// fmt.Println(dp3)

	for i := 0; i < 4; i++ {
		mask1 := 1 << i
		a := float32(nums[i])
		res := dp3[mask1^fullMask]
		for b := range res {
			// + - *
			if near24(a+b) || near24(a-b) || near24(b-a) || near24(a*b) {
				// fmt.Println(a, b)
				return true
			}
			// /
			if b != 0 && (near24(a/b) || near24(b/a)) {
				// fmt.Println(a, b)
				return true
			}
		}
	}
	return false
}

func near24(num float32) bool {
	num -= 24
	if num > precision || num < -precision {
		return false
	}
	return true
}

// =============== brute force helper ===============

type ch byte

func (c ch) isOp() bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func (c ch) value() *big.Rat {
	return big.NewRat(int64(c-'0'), 1)
}

func (c ch) apply(a, b *big.Rat) *big.Rat {
	res := new(big.Rat)
	switch c {
	case '+':
		return res.Add(a, b)
	case '-':
		return res.Sub(a, b)
	case '*':
		return res.Mul(a, b)
	case '/':
		return res.Quo(a, b)
	}
	return res
}

var (
	ops    = [...]ch{'+', '-', '*', '/'}
	nn     []int
	used   []bool
	exp    = make([]ch, 7)
	opLeft int
	exps   [][]ch
)

func bruteForce(nums []int) bool {
	nn = nums
	used = make([]bool, 4)
	exps = make([][]ch, 0)
	opLeft = 2

	for i := 0; i < 4; i++ {
		exp[0] = ops[i]
		chDfs(1)
	}
	for _, e := range exps {
		if is24(e) {
			buf := make([]byte, 7)
			for i := 0; i < 7; i++ {
				buf[i] = byte(e[i])
			}
			//fmt.Println(string(buf))
			return true
		}
	}
	return false
}

func chDfs(i int) {
	if opLeft > 0 {
		opLeft--
		for j := 0; j < 4; j++ {
			exp[i] = ops[j]
			chDfs(i + 1)
		}
		opLeft++
	}

	if i+opLeft < 5 || opLeft == 0 {
		for j := 0; j < 4; j++ {
			if used[j] {
				continue
			}

			exp[i] = '0' + ch(nn[j])
			if i == 6 {
				c := make([]ch, 7)
				copy(c, exp)
				exps = append(exps, c)
			} else {
				used[j] = true
				chDfs(i + 1)
				used[j] = false
			}
		}
	}
}

func is24(e []ch) bool {
	defer func() {
		recover()
	}()
	res, ed := calc(e, 0)
	if ed == 7 && res.IsInt() && res.Num().Int64() == 24 {
		return true
	}
	return false
}

func calc(e []ch, st int) (*big.Rat, int) {
	var a, b *big.Rat

	op := e[st]
	st++

	cha := e[st]
	if cha.isOp() {
		a, st = calc(e, st)
	} else {
		a = cha.value()
		st++
	}

	chb := e[st]
	if chb.isOp() {
		b, st = calc(e, st)
	} else {
		b = chb.value()
		st++
	}

	return op.apply(a, b), st
}

// func judgePoint24(nums []int) bool {
// 	sort.Ints(nums)
// 	var dp = map[int](map[float32]bool){}

// 	for i := 0; i < 3; i++ {
// 		for j := i + 1; j < 4; j++ {
// 			id := nums[i]*10 + nums[j]
// 			if _, ok := dp[id]; ok {
// 				continue
// 			}
// 			a := float32(nums[i])
// 			b := float32(nums[j])
// 			result := make(map[float32]bool)
// 			// +
// 			result[a+b] = true
// 			// -
// 			result[a-b] = true
// 			result[b-a] = true
// 			// *
// 			result[a*b] = true
// 			// /
// 			result[a/b] = true
// 			result[b/a] = true

// 			dp[id] = result
// 		}
// 	}

// 	for id1, res1 := range dp {
// 		for id2, res2 := range dp {
// 			if id1&id2 != 0 {
// 				continue
// 			}

// 			for a := range res1 {
// 				for b := range res2 {
// 					// + - *
// 					if a+b == 24 || a-b == 24 || b-a == 24 || a*b == 24 {
// 						return true
// 					}
// 					// /
// 					if a != 0 && b != 0 && (a/b == 24 || b/a == 24) {
// 						return true
// 					}
// 				}
// 			}
// 		}
// 	}

// 	for i := 0; i < 4; i++ {
// 		id1 := 1 << i
// 		a := float32(nums[i])
// 		for id2, res := range dp {
// 			if id1&id2 != 0 {
// 				continue
// 			}
// 			id := id1 | id2
// 			result := make(map[float32]bool)
// 			for b := range res {
// 				// + - *
// 				result[a+b] = true
// 				result[a-b] = true
// 				result[b-a] = true
// 				result[a*b] = true
// 				result[b/a] = true
// 				if b != 0 {
// 					result[a/b] = true
// 				}
// 			}
// 			dp[id] = result
// 		}
// 	}

// 	for i := 0; i < 4; i++ {
// 		id1 := 1 << i
// 		a := float32(nums[i])
// 		for id2, res := range dp {
// 			if id1|id2 != fullMask {
// 				continue
// 			}
// 			for b := range res {
// 				// + - *
// 				if a+b == 24 || a-b == 24 || b-a == 24 || a*b == 24 {
// 					return true
// 				}
// 				// /
// 				if b != 0 && (a/b == 24 || b/a == 24) {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }
