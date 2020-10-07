package main

// 12ms
// 5.8MB
func largestRectangleArea(heights []int) int {
	maxArea := 0
	s := make(stack, 1)
	s[0] = &bar{ // sentry
		right:  0,
		height: 0,
	}
	for i, h := range heights {
		for len(s) > 1 && h < s.Top().height {
			b := s.Pop()
			area := (i - s.Top().right) * b.height
			if maxArea < area {
				maxArea = area
			}
		}
		top := s.Top()
		if h > top.height { // avoid ==
			s.Push(&bar{
				right:  i + 1,
				height: h,
			})
		} else if h == top.height {
			top.right = i + 1
		}
	}
	width := len(heights)
	for len(s) > 1 {
		b := s.Pop()
		area := (width - s.Top().right) * b.height
		if maxArea < area {
			maxArea = area
		}
	}
	return maxArea
}

type bar struct {
	right  int
	height int
}

type stack []*bar

func (s stack) Top() *bar      { return s[len(s)-1] }
func (s *stack) Push(val *bar) { *s = append(*s, val) }
func (s *stack) Pop() *bar {
	old := *s
	newSize := len(old) - 1
	*s = old[:newSize]
	return old[newSize]
}
