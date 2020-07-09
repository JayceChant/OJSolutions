package main

import (
	"container/list"
	"sort"
	"strings"
)

// k 次交换获得最小整数
// 1 <= num.length <= 30000
// num 只包含 数字 且不含有 前导 0 。
// 1 <= k <= 10^9

// 超时
func minIntegerBF(num string, k int) string {
	digits := []byte(num)
	for target := 0; target < len(digits) && k > 0; target++ {
		from := target
		for i := target + 1; i < len(digits) && i-target <= k; i++ {
			if digits[i] < digits[from] {
				from = i
			}
		}
		if from == target {
			continue
		}
		tmp := digits[from]
		for i := from; i > target; i-- {
			digits[i] = digits[i-1]
		}
		digits[target] = tmp
		k -= (from - target)
	}
	return string(digits)
}

// 更慢，失败的优化
func minIntegerBF2(num string, k int) string {
	digits := []byte(num)
	sorted := make([]byte, len(digits))
	copy(sorted, digits)
	sort.Sort(byteSlice(sorted))
	target := 0
	for minIdx := 0; k > 0 && minIdx < len(sorted); minIdx++ {
		from := target + 1
		for ; from < len(digits); from++ {
			if digits[from] == sorted[minIdx] {
				break
			}
		}
		if from == len(digits) {
			target++
			continue
		}
		offset := from - target
		if offset > k {
			continue
		}
		tmp := digits[from]
		for i := from; i > target; i-- {
			digits[i] = digits[i-1]
		}
		digits[target] = tmp
		target++
		k -= offset
	}
	return string(digits)
}

var _ sort.Interface = byteSlice(nil)

type byteSlice []byte

func (s byteSlice) Len() int {
	return len(s)
}

func (s byteSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s byteSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func minIntegerRec(num string, k int) string {
	sb := new(strings.Builder)
	return minIntegerRecImpl(sb, []byte(num), k).String()
}

func minIntegerRecImpl(builder *strings.Builder, digits []byte, k int) *strings.Builder {
	if len(digits) == 0 {
		return builder
	}

	if k == 0 {
		builder.Write(digits)
		return builder
	}

	minIdx := 0

	for i := 0; i <= k && i < len(digits); i++ {
		if digits[i] < digits[minIdx] {
			minIdx = i
		}
	}

	builder.WriteByte(digits[minIdx])

	return minIntegerRecImpl(builder, append(digits[:minIdx], digits[minIdx+1:]...), k-minIdx)
}

// cn: 144 ms , 在所有 Go 提交中击败了 76.00% 的用户 (min 12ms)
func minIntegerST(num string, k int) string {
	digits := []byte(num)

	var idxQueues [10]*list.List
	for i := 0; i < 10; i++ {
		idxQueues[i] = list.New()
	}

	for i, d := range digits {
		idxQueues[d-byte('0')].PushBack(i)
	}

	st := buildSegmentTree(len(digits))
	//fmt.Println(st)

	ans := strings.Builder{}
	for to := 0; to < len(digits); to++ {
		for digit := byte(0); digit <= 9; digit++ {
			if idxQueues[digit].Len() == 0 {
				continue
			}
			e := idxQueues[digit].Front()
			from := e.Value.(int)
			move := st.query(from+1, len(digits)-1)
			offset := from + move - to
			//fmt.Println("try", to, from, move)
			if offset <= k {
				ans.WriteByte(byte('0') + digit)
				k -= offset
				idxQueues[digit].Remove(e)
				st.increase(from)
				break
			}
		}
	}
	return ans.String()
}

type treeNode struct {
	Sum   int
	Start int
	End   int
	Left  *treeNode
	Right *treeNode
}

func buildSegmentTree(size int) *treeNode {
	//fmt.Println("buildTree", size)
	root := buildNode(0, size-1)
	return root
}

func buildNode(start int, end int) *treeNode {
	//fmt.Println("build", start, end)
	root := &treeNode{Start: start, End: end}
	if start == end {
		return root
	}
	mid := (start + end) / 2
	root.Left = buildNode(start, mid)
	root.Right = buildNode(mid+1, end)
	return root
}

func (root *treeNode) increase(index int) {
	// fmt.Println("increase", index, root.Start, root.End)
	root.Sum++
	if root.Start == root.End {
		return
	}
	mid := (root.Start + root.End) / 2
	if index <= mid {
		root.Left.increase(index)
	} else {
		root.Right.increase(index)
	}
}

func (root *treeNode) query(left int, right int) int {
	if left > right {
		return 0
	}
	//fmt.Println("getSum", left, right, root.Start, root.End)
	if left <= root.Start && root.End <= right {
		return root.Sum
	}
	mid := (root.Start + root.End) / 2
	sum := 0
	if left <= mid {
		sum += root.Left.query(left, right)
	}
	if right > mid {
		sum += root.Right.query(left, right)
	}
	return sum
}

func minIntegerBIT(num string, k int) string {
	digits := []byte(num)

	var idxQueues [10]*list.List
	for i := 0; i < 10; i++ {
		idxQueues[i] = list.New()
	}

	for i, d := range digits {
		idxQueues[d-byte('0')].PushBack(i)
	}

	bit := make(binaryIndexedTree, len(digits)+1)
	//fmt.Println(num)

	ans := strings.Builder{}
	for to := 0; to < len(digits); to++ {
		for digit := byte(0); digit <= 9; digit++ {
			if idxQueues[digit].Len() == 0 {
				continue
			}
			e := idxQueues[digit].Front()
			from := e.Value.(int)
			move := bit.queryRange(from+2, len(digits))
			offset := from + move - to
			//fmt.Println("try", to, from, move)
			if offset <= k {
				ans.WriteByte(byte('0') + digit)
				k -= offset
				idxQueues[digit].Remove(e)
				bit.update(from+1, 1)
				break
			}
		}
	}
	return ans.String()
}

type binaryIndexedTree []int

func (bit binaryIndexedTree) queryTo(index int) int {
	//fmt.Println("queryTo", index)
	sum := 0
	for ; index > 0; index -= lowbit(index) {
		sum += bit[index]
	}
	return sum
}

func (bit binaryIndexedTree) queryRange(left int, right int) int {
	//fmt.Println("queryRange", left, right)
	return bit.queryTo(right) - bit.queryTo(left-1)
}

func (bit binaryIndexedTree) update(index int, change int) {
	//fmt.Println("update", index, change)
	if index <= 0 {
		return
	}

	for ; index < len(bit); index += lowbit(index) {
		bit[index] += change
	}
}

func lowbit(x int) int {
	return x & (-x)
}
