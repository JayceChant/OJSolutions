package main

import (
	"container/heap"
)

// ListNode 题目运行环境提供，实际提交时不需要
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKListsOneByOne(lists)
}

func mergeKListsOneByOne(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	l := lists[0]
	for i := 1; i < len(lists); i++ {
		l = merge(l, lists[i])
	}
	return l
}

func merge(la *ListNode, lb *ListNode) *ListNode {
	head := new(ListNode)
	tail := head
	for la != nil && lb != nil {
		if la.Val < lb.Val {
			tail.Next = la
			la = la.Next
		} else {
			tail.Next = lb
			lb = lb.Next
		}
		tail = tail.Next
	}

	if la == nil {
		tail.Next = lb
	} else {
		tail.Next = la
	}

	return head.Next
}

func mergeKListsInsert(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	l := lists[0]
	for i := 1; i < len(lists); i++ {
		l = insert(l, lists[i])
	}
	return l
}

func insert(la *ListNode, lb *ListNode) *ListNode {
	head := new(ListNode)
	head.Next = la
	tail := head
	for tail.Next != nil && lb != nil {
		if lb.Val < tail.Next.Val {
			tail.Next, lb.Next, lb = lb, tail.Next, lb.Next
		}
		tail = tail.Next
	}
	if tail.Next == nil {
		tail.Next = lb
	}
	return head.Next
}

// DAC = Divide and Conquer
func mergeKListsInsertDAC(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	k := len(lists)

	for step := 1; step < k; {
		width := step * 2
		for i := 0; i < len(lists); i += width {
			if i+step < k {
				lists[i] = insert(lists[i], lists[i+step])
			}
		}
		step = width
	}
	return lists[0]
}

type myHeap []*ListNode

func (h *myHeap) push(node *ListNode) {
	cur := len(*h)
	*h = append(*h, node)
	hh := *h
	// bubble up
	for cur > 0 {
		up := (cur - 1) / 2
		if hh[up].Val <= hh[cur].Val {
			break
		}
		hh[up], hh[cur] = hh[cur], hh[up]
		cur = up
	}
}

func (h *myHeap) pop() *ListNode {
	hh := *h
	node := hh[0]
	n := len(hh)
	hh[0] = hh[n-1]
	// bubble down
	i := 0
	for {
		down := i*2 + 1
		if down >= n {
			break
		}
		if down+1 < n && hh[down].Val > hh[down+1].Val {
			down++
		}
		if hh[i].Val <= hh[down].Val {
			break
		}
		hh[i], hh[down] = hh[down], hh[i]
		i = down
	}
	hh[n-1] = nil // 避免内存泄漏
	*h = hh[:n-1]
	return node
}

func mergeKListsWithHeap(lists []*ListNode) *ListNode {
	h := make(myHeap, 0)
	for _, l := range lists {
		for l != nil {
			h.push(l)
			l = l.Next
		}
	}

	head := new(ListNode)
	tail := head
	for len(h) != 0 {
		tail.Next = h.pop()
		tail = tail.Next
	}
	return head.Next
}

func mergeKListsWithHeap2(lists []*ListNode) *ListNode {
	h := make(myHeap, 0)
	for _, l := range lists {
		if l != nil {
			h.push(l)
		}
	}

	head := new(ListNode)
	tail := head
	for len(h) != 0 {
		tail.Next = h.pop()
		tail = tail.Next
		if tail.Next != nil {
			h.push(tail.Next)
		}
	}
	return head.Next
}

func (h myHeap) top() *ListNode {
	return h[0]
}

func (h myHeap) update(i int, node *ListNode) {
	h[i] = node
	n := len(h)
	for {
		down := i*2 + 1
		if down >= n {
			break
		}
		if down+1 < n && h[down].Val > h[down+1].Val {
			down++
		}
		if h[i].Val <= h[down].Val {
			break
		}
		h[i], h[down] = h[down], h[i]
		i = down
	}
}

func (h *myHeap) remove(i int) {
	hh := *h
	n := len(hh)
	hh[i] = hh[n-1]
	// bubble down
	for {
		down := i*2 + 1
		if down >= n {
			break
		}
		if down+1 < n && hh[down].Val > hh[down+1].Val {
			down++
		}
		if hh[i].Val <= hh[down].Val {
			break
		}
		hh[i], hh[down] = hh[down], hh[i]
		i = down
	}
	hh[n-1] = nil
	*h = hh[:n-1]
}

func mergeKListsWithHeap3(lists []*ListNode) *ListNode {
	h := make(myHeap, 0)
	for _, l := range lists {
		if l != nil {
			h.push(l)
		}
	}

	head := new(ListNode)
	tail := head
	for len(h) != 0 {
		tail.Next = h.top()
		tail = tail.Next
		if tail.Next != nil {
			h.update(0, tail.Next)
		} else {
			h.remove(0)
		}
	}
	return head.Next
}

var _ heap.Interface = (*nodeHeap)(nil)

type nodeHeap []*ListNode

func (h nodeHeap) Len() int { return len(h) }

func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h nodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }

func (h *nodeHeap) Pop() interface{} {
	n := len(*h) - 1
	node := (*h)[n]
	*h = (*h)[:n]
	return node
}

func mergeKListsWithBuiltinHeap(lists []*ListNode) *ListNode {
	h := make(nodeHeap, 0)
	for _, l := range lists {
		if l != nil {
			heap.Push(&h, l)
		}
	}

	head := new(ListNode)
	tail := head
	for len(h) != 0 {
		tail.Next = h[0]
		tail = tail.Next
		if tail.Next != nil {
			h[0] = tail.Next
			heap.Fix(&h, 0)
		} else {
			heap.Remove(&h, 0)
		}
	}
	return head.Next
}

// debug only

// func (h *heap) String() string {
// 	sb := strings.Builder{}
// 	sb.WriteRune('[')
// 	for _, n := range *h {
// 		sb.WriteString(fmt.Sprintf("%d,", n.Val))
// 	}
// 	sb.WriteRune(']')
// 	return sb.String()
// }
