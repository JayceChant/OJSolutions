package main

// 232 ms, 在所有 Go 提交中击败了96.89%的用户
// 8.2 MB, 在所有 Go 提交中击败了100.00%的用户
type CQueue struct {
	inStack  stack
	outStack stack
}

func Constructor() CQueue {
	return CQueue{
		inStack:  make(stack, 0),
		outStack: make(stack, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack.push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.outStack.len() > 0 {
		return this.outStack.pop()
	}

	if this.inStack.len() == 0 {
		return -1
	}

	for this.inStack.len() > 1 {
		this.outStack.push(this.inStack.pop())
	}
	return this.inStack.pop()
}

type stack []int

func (s stack) len() int {
	return len(s)
}

func (s *stack) push(e int) {
	*s = append(*s, e)
}

func (s *stack) pop() int {
	tail := len(*s) - 1
	e := (*s)[tail]
	*s = (*s)[:tail]
	return e
}
