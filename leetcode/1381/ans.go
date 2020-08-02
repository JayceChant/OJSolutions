package main

// CustomStack ...
// use container/list
// 24ms
// 6.6MB
// use []int
// 24ms
// 6.8MB
type CustomStack struct {
	data []int
	cap  int
	inc  []int
}

// Constructor ...
func Constructor(maxSize int) CustomStack {
	return CustomStack{
		data: make([]int, 0, maxSize),
		cap:  maxSize,
		inc:  make([]int, maxSize),
	}
}

// Push ...
func (s *CustomStack) Push(x int) {
	if len(s.data) < s.cap {
		s.data = append(s.data, x)
	}
}

// Pop ...
func (s *CustomStack) Pop() int {
	k := len(s.data) - 1 // start from 0
	if k == -1 {
		return -1
	}

	val := s.data[k] + s.inc[k]
	s.data = s.data[:k]
	if k > 0 {
		s.inc[k-1] += s.inc[k]
	}
	s.inc[k] = 0
	return val
}

// Increment ...
func (s *CustomStack) Increment(k int, val int) {
	if k > len(s.data) {
		k = len(s.data)
	}
	if k > 0 {
		s.inc[k-1] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
