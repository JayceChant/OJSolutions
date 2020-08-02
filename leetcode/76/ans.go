package main

// 12ms, 84.46% (min 0ms)
// 3.4MB, 52%
func minWindowQueue(src string, tar string) string {
	letterQues := map[byte](*queue){}
	for i := range tar {
		if letterQues[tar[i]] == nil {
			letterQues[tar[i]] = newQueue()
		}
		letterQues[tar[i]].cap++
	}

	matchQue := newDeque()
	tLen := len(tar)
	min := len(src)
	left := 0
	right := -1
	included := false

	for i := range src {
		ch := src[i]
		if letterQues[ch] == nil {
			continue
		}

		n, full := letterQues[ch].pushBack(i)
		movedHead := matchQue.pushBack(n, full)

		if !included {
			if matchQue.size < tLen {
				continue
			}
			included = true
		} else {
			if !movedHead {
				continue
			}
		}

		diff := i - matchQue.front().index
		if min > diff {
			min = diff
			left = matchQue.front().index
			right = i
		}
	}

	return src[left : right+1]
}

var (
	lastIndex map[byte](*node)
)

type node struct {
	index    int
	nextSame *node
	prev     *node
	next     *node
}

type queue struct {
	sentry node
	tail   *node
	size   int
	cap    int
}

func newQueue() *queue {
	q := new(queue)
	q.tail = &q.sentry
	return q
}

// if the capacity is full, pop the front node, change the value, and push to back again
// return the node just added or moved
func (q *queue) pushBack(index int) (n *node, full bool) {
	if q.size < q.cap {
		n = &node{index: index}
		q.tail.nextSame = n
		q.tail = n
		q.size++
		return
	}
	full = true
	n = q.sentry.nextSame
	// reuse
	n.index = index
	if q.size == 1 {
		return
	}
	// pop head
	q.sentry.nextSame = n.nextSame
	n.nextSame = nil
	// push back
	q.tail.nextSame = n
	q.tail = n
	return
}

type deque struct {
	root node
	size int
}

func newDeque() *deque {
	q := new(deque)
	sentry := &q.root
	q.root.prev = sentry
	q.root.next = sentry
	return q
}

func (q *deque) pushBack(n *node, full bool) (isHead bool) {
	if full {
		isHead = (n == q.root.next)
		n.prev.next = n.next
		n.next.prev = n.prev
	}
	tail := q.root.prev
	n.prev = tail
	n.next = tail.next
	tail.next = n
	q.root.prev = n
	if !full {
		q.size++
	}
	return
}

func (q *deque) front() *node {
	return q.root.next
}

// sliding window
// 24ms
// 2.9MB
func minWindow(src string, tar string) string {
	sLen := len(src)
	if sLen == 0 || len(tar) == 0 {
		return ""
	}

	requiredCounts := map[byte]int{}
	for i := range tar {
		requiredCounts[tar[i]]++
	}
	requiredLetters := len(requiredCounts)

	min := sLen + 1
	var minLeft, minRight, left, right int

	for ; right < sLen && requiredLetters > 0; right++ {
		ch := src[right]
		if _, inTar := requiredCounts[ch]; !inTar {
			continue
		}

		requiredCounts[ch]--
		if requiredCounts[ch] == 0 {
			requiredLetters--
		}
	}

	if requiredLetters > 0 {
		return ""
	}

	for { // right < sLen
		for ; ; left++ {
			ch := src[left]
			if _, inTar := requiredCounts[ch]; !inTar {
				continue
			}

			requiredCounts[ch]++
			if requiredCounts[ch] > 0 {
				break
			}
		}

		width := right - left
		if min > width {
			min = width
			minLeft = left
			minRight = right
		}

		if right == sLen {
			break
		}

		lackCh := src[left]
		for ; right < sLen && requiredCounts[lackCh] > 0; right++ {
			ch := src[right]
			if _, inTar := requiredCounts[ch]; !inTar {
				continue
			}
			requiredCounts[ch]--
		}

		if requiredCounts[lackCh] > 0 {
			break
		}
		left++
	}

	return src[minLeft:minRight]
}
