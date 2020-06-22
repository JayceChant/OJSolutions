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
	return mergeKListsInsertDAC(lists)
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
		if lb.Val <= tail.Next.Val {
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
	k := len(lists)
	if k == 0 {
		return nil
	}

	// 空白伪头结点，避免判空处理
	head := new(ListNode)
	// 合并步长从 1 开始不断翻倍
	for step := 1; step < k; {
		// 分组宽度是步长的两倍
		width := step * 2
		for i := 0; i+step < k; i += width {
			// func merge(list1 *ListNode, list2 *ListNode) 实际开发中这部分应该抽取子函数
			head.Next = lists[i]
			list2 := lists[i+step]
			tail := head
			for tail.Next != nil && list2 != nil {
				if list2.Val <= tail.Next.Val {
					tail.Next, list2.Next, list2 = list2, tail.Next, list2.Next
				}
				tail = tail.Next
			}
			if list2 != nil {
				tail.Next = list2
			}
			// 第一个节点可能是来自 list2，合并结果一定要赋值回去
			lists[i] = head.Next
			// func merge() end
		}
		step = width
	}
	return lists[0]
}

type myHeap []*ListNode

func (hp *myHeap) push(node *ListNode) {
	child := len(*hp)
	*hp = append(*hp, node)
	h := *hp
	// bubble up
	for child > 0 {
		parent := (child - 1) / 2
		if h[parent].Val <= h[child].Val {
			break
		}
		h[parent], h[child] = h[child], h[parent]
		child = parent
	}
}

func (hp *myHeap) pop() *ListNode {
	h := *hp
	top := h[0]
	k := len(h)
	h[0] = h[k-1]
	// bubble down
	parent := 0
	for {
		child := parent*2 + 1
		if child >= k {
			break
		}
		if child+1 < k && h[child].Val > h[child+1].Val {
			child++
		}
		if h[parent].Val <= h[child].Val {
			break
		}
		h[parent], h[child] = h[child], h[parent]
		parent = child
	}
	h[k-1] = nil // 避免内存泄漏
	*hp = h[:k-1]
	return top
}

func mergeKListsWithNHeap(lists []*ListNode) *ListNode {
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

func mergeKListsWithHeapPushPop(lists []*ListNode) *ListNode {
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

func (h myHeap) build() {
	k := len(h)
	for parent := k/2 - 1; parent >= 0; parent-- {
		child := parent*2 + 1
		if child+1 < k && h[child].Val > h[child+1].Val {
			child++
		}
		if h[parent].Val <= h[child].Val {
			continue
		}
		h[parent], h[child] = h[child], h[parent]
	}
}

func (h myHeap) top() *ListNode {
	return h[0]
}

func (h myHeap) update(i int, node *ListNode) {
	h[i] = node
	k := len(h)
	for {
		child := i*2 + 1
		if child >= k {
			break
		}
		if child+1 < k && h[child].Val > h[child+1].Val {
			child++
		}
		if h[i].Val <= h[child].Val {
			break
		}
		h[i], h[child] = h[child], h[i]
		i = child
	}
}

func (hp *myHeap) remove(i int) {
	h := *hp
	k := len(h)
	h[i] = h[k-1]
	// bubble down
	for {
		child := i*2 + 1
		if child >= k {
			break
		}
		if child+1 < k && h[child].Val > h[child+1].Val {
			child++
		}
		if h[i].Val <= h[child].Val {
			break
		}
		h[i], h[child] = h[child], h[i]
		i = child
	}
	h[k-1] = nil
	*hp = h[:k-1]
}

func mergeKListsWithHeapUpdateRemove(lists []*ListNode) *ListNode {
	h := make(myHeap, 0)
	for _, l := range lists {
		if l != nil {
			h = append(h, l)
		}
	}
	h.build()

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
			h = append(h, l)
		}
	}
	heap.Init(&h)

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
