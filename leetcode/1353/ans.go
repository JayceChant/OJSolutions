package main

import (
	"container/heap"
	"sort"
)

// 320ms (min 180ms)
// 17.9MB
func maxEventsSlideSortedWindow(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	//fmt.Println(events)

	count := 0
	var st, ed int
	day := 1 // the 1st free day
	size := len(events)
	for st < size {
		extend := false
		for ed < size {
			for ed < size && events[ed][0] <= day {
				ed++ // include all events start at day
				extend = true
			}
			if extend {
				sort.Slice(events[st:ed], func(i, j int) bool {
					return events[st+i][1] < events[st+j][1]
				})
				//fmt.Println("day=", day, ",[", st, ed, "),", events[st:ed])
				break
			}
			if ed-st > 0 {
				break
			}
			day = events[ed][0]
		}

		for st < size && events[st][1] < day {
			// skip the events already ended
			st++
		}

		if st == size {
			break
		}

		if st == ed {
			continue
		}

		if day < events[st][0] {
			day = events[st][0]
		}
		count++
		day++
		st++
	}
	return count
}

// 352ms
// 31.8MB
func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	count := 0
	day := 1 // the 1st free day
	size := len(events)
	i := 0
	h := make(minEnd, 0)
	for i < size || len(h) > 0 {
		for i < size {
			for i < size && events[i][0] <= day {
				// include all events already started
				heap.Push(&h, events[i])
				i++
			}
			if len(h) > 0 {
				break
			}
			day = events[i][0]
		}

		//fmt.Println(h)

		for len(h) > 0 && h[0][1] < day {
			// skip the events already ended
			heap.Pop(&h)
		}

		if len(h) == 0 {
			continue
		}

		//fmt.Println(h)

		count++
		day++
		heap.Pop(&h)
	}
	return count
}

type minEnd [][]int

var _ heap.Interface = (*minEnd)(nil)

func (a minEnd) Len() int            { return len(a) }
func (a minEnd) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a minEnd) Less(i, j int) bool  { return a[i][1] < a[j][1] }
func (a *minEnd) Push(x interface{}) { *a = append(*a, x.([]int)) }
func (a *minEnd) Pop() interface{} {
	old := *a
	newSize := len(old) - 1
	*a = old[:newSize]
	return old[newSize]
}
