package main

// SummaryRanges ...
// 56ms
// 12.4MB
type SummaryRanges struct {
	*node
}

// Constructor ...
/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{new(node)}
}

// AddNum ...
func (r *SummaryRanges) AddNum(val int) {
	slow := r.node
	fast := r.next
	left := val - 1
	for fast != nil && fast.rang[1] < left {
		slow = fast
		fast = fast.next
	}

	if fast == nil || val < fast.rang[0]-1 {
		slow.next = &node{
			rang: []int{val, val},
			next: fast,
		}
		return
	}

	if fast.rang[1] == left {
		// append to fast
		fast.rang[1] = val

		slow = fast
		fast = fast.next
		if fast == nil || fast.rang[0] > val+1 {
			return
		}
		// connected, merge
		slow.rang[1] = fast.rang[1]
		slow.next = fast.next
		fast.next = nil
		return
	}

	if fast.rang[0] <= val && val <= fast.rang[1] {
		// val is included in fast
		return
	}

	//if right := val + 1; right == fast.rang[0] {
	fast.rang[0] = val
	//	return
	//}
}

// GetIntervals ...
func (r *SummaryRanges) GetIntervals() (ans [][]int) {
	tail := r.next
	for tail != nil {
		ans = append(ans, tail.rang)
		tail = tail.next
	}
	return
}

type node struct {
	rang []int
	next *node
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
