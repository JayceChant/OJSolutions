package main

import (
	"math"
)

func removeBoxes(boxes []int) int {
	colors = []int{boxes[0]}
	counts = []int{1}
	total = len(boxes)

	ci := 0
	for i := 1; i < total; i++ {
		if boxes[i] != colors[ci] {
			colors = append(colors, boxes[i])
			counts = append(counts, 1)
			ci++
			continue
		}
		counts[ci]++
	}
	total = len(colors)

	universe = newUniverse(total)

	removed.clearAll()
	cache = make(map[bitmap128]int)
	removeCount = 0
	precalc()
	// fmt.Println(total)
	// fmt.Println(cache)
	// fmt.Println(combo)
	dp()
	//fmt.Println(cache)
	return cache[*universe]
}

var (
	colors, counts     []int
	removeCount, total int
	universe           *bitmap128
	removed            bitmap128
	combo              map[int][]bitmap128
	cache              map[bitmap128]int
)

func precalc() {
	combo = make(map[int][]bitmap128)
	for st := 0; st < total; st++ {
		removed.set(st)
		length := counts[st]
		cache[removed] = length * length
		dfs(colors[st], st, st+2, length)
		removed.clear(st)
	}
}

func dfs(color int, st int, i int, length int) {
	for i < total && colors[i] != color {
		i++
	}
	if i >= total {
		return
	}

	// skip
	dfs(color, st, i+2, length)

	removed.set(i)
	length += counts[i]
	cache[removed] = length * length
	id := (i-st+1)*100 + st
	combo[id] = append(combo[id], removed)
	dfs(color, st, i+2, length)
	removed.clear(i)
}

func dp() {
	var left bitmap128
	for width := 2; width <= total; width++ {
		//fmt.Println(width)
		u := newUniverse(width)
		// [st, ed)
		for st := 0; st <= total-width; st++ {
			if cbs, ok := combo[width*100+st]; ok {
				for _, cb := range cbs {
					sum := cache[cb]
					for _, g := range cb.complement(u).groups() {
						sum += cache[*g]
					}
					if cache[*u] < sum {
						cache[*u] = sum
					}
				}
			}

			ed := st + width
			left.clearAll()
			for i := st; i < ed; i++ {
				left.set(i)
				right := left.complement(u)
				sum := cache[left] + cache[*right]
				if cache[*u] < sum {
					// fmt.Println(left)
					cache[*u] = sum
				}
			}
			u.lshift()
		}
	}
}

type bitmap128 [2]uint64

func newUniverse(n int) *bitmap128 {
	if n <= 64 {
		return &bitmap128{math.MaxUint64 >> (64 - n), 0}
	}
	return &bitmap128{
		math.MaxUint64,
		math.MaxUint64 >> (128 - n),
	}
}

// func (b *bitmap128) get(i int) bool { return (*b)[i/64]&(1<<(i%64)) != 0 }
func (b *bitmap128) set(i int)   { (*b)[i/64] |= (1 << (i % 64)) }
func (b *bitmap128) clear(i int) { (*b)[i/64] &= (^(1 << (i % 64))) }

// func (b *bitmap128) union(b2 *bitmap128) {
// 	(*b)[0] |= (*b2)[0]
// 	(*b)[1] |= (*b2)[1]
// }
func (b *bitmap128) minus(b2 *bitmap128) {
	(*b)[0] &= (^(*b2)[0])
	(*b)[1] &= (^(*b2)[1])
}
func (b *bitmap128) lshift() {
	(*b)[1] = ((*b)[1] << 1) | ((*b)[0] >> 63)
	(*b)[0] <<= 1
}
func (b *bitmap128) complement(u *bitmap128) *bitmap128 {
	c := *u
	c[0] ^= (*b)[0]
	c[1] ^= (*b)[1]
	return &c
}
func (b *bitmap128) groups() []*bitmap128 {
	g := make([]*bitmap128, 0)
	bc := *b // b clone
	for bc[0] > 0 {
		c := bc // clone of bc
		lowbit := bc[0] & (^bc[0] + 1)
		bc[0] += lowbit
		lowbit = bc[0] & (^bc[0] + 1)
		bc[0] -= lowbit
		if bc[0] == 0 && c[0]>>63 == 1 && bc[1]&1 == 1 {
			bc[1]++
			lowbit = bc[1] & (^bc[1] + 1)
			bc[1] -= lowbit
		}
		c.minus(&bc)
		g = append(g, &c)
	}
	for bc[1] > 0 {
		c := bc
		lowbit := bc[1] & (^bc[1] + 1)
		bc[1] += lowbit
		lowbit = bc[1] & (^bc[1] + 1)
		bc[1] -= lowbit
		c.minus(&bc)
		g = append(g, &c)
	}
	return g
}
func (b *bitmap128) clearAll() {
	(*b)[0] = 0
	(*b)[1] = 0
}

// func (b bitmap128) String() string {
// 	return fmt.Sprintf("%b-%064b", b[1], b[0])
// }
