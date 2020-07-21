package main

import (
	"sort"
)

func maxNumOfSubstrings(s string) []string {
	// segsByLetter[i] indicates the segment of string that inlcude ('a' + i)
	segsByLetter := make([]*segment, 26)
	for i := 0; i < 26; i++ {
		// pre-init to avoid NPE
		segsByLetter[i] = new(segment)
	}

	segsByLeftDesc := make(segsSortByLeftDesc, 26)
	// share same segment nodes but in different order
	copy(segsByLeftDesc, segsByLetter)

	size := len(s)
	left := 0
	right := 0
	// [left, right)
	for {
		left = right
		if left >= size {
			break
		}

		right++
		for right < size && s[right] == s[right-1] {
			right++
		}
		i := s[left] - 'a'

		// empty segment for current character
		if segsByLetter[i].right == 0 {
			segsByLetter[i].left = left
			segsByLetter[i].right = right
			continue
		}

		//fmt.Printf("%d:%c\n", right, s[left])

		// segment for current character exists, try merge
		seg := segsByLetter[i]
		for {
			oldLeft := seg.left
			seg.right = right

			merged := false
			sort.Sort(segsByLeftDesc)
			for _, segInter := range segsByLeftDesc {
				if segInter == seg || segInter.right == 0 {
					continue
				}

				// check and update intersection, make them equivalent segments
				// the segments intersected with current segment are equivalents,
				// because they are intersected with each other and merged earlier

				// merge 1 segment only, to avoid skipping update the equivalents
				if !merged && segInter.left < seg.left && segInter.right > seg.left {
					// merge the first intersection encountered
					seg.left = segInter.left
					segInter.right = seg.right
					merged = true
					continue
				}

				if segInter.left == seg.left && segInter.right < seg.right {
					// update right of rest segments equivalent to the merged one
					segInter.right = seg.right
				}
				//printSegments(segsByLetter)
			}

			if !merged {
				break
			}

			for _, segEq := range segsByLeftDesc {
				if segEq == seg || segEq.right == 0 {
					continue
				}

				if segEq.left < oldLeft {
					break
				}

				if segEq.left == oldLeft {
					// update left of rest segments equivalent to seg
					segEq.left = seg.left
					segEq.right = seg.right
				}
			}
		} // for merge all intersected segments
	} // for transverse string

	//printSegments(segsByLetter)
	sort.Sort(segsSortByLeftAndLen(segsByLetter))
	ans := make([]string, 0)
	minLeft := 0
	for _, seg := range segsByLetter {
		if seg.right == 0 || seg.left < minLeft {
			continue
		}
		ans = append(ans, s[seg.left:seg.right])
		minLeft = seg.right
	}
	return ans
}

type segment struct {
	// [left, right)
	left  int
	right int
}

var _ sort.Interface = segsSortByLeftAndLen(nil)
var _ sort.Interface = segsSortByLeftDesc(nil)

type segsSortByLeftAndLen []*segment

func (a segsSortByLeftAndLen) Len() int      { return len(a) }
func (a segsSortByLeftAndLen) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a segsSortByLeftAndLen) Less(i, j int) bool {
	return a[i].right <= a[j].left || // i on the left of j with no intersection
		// or i is shorter with intersection
		(a[i].left < a[j].right && (a[i].right-a[i].left) < (a[j].right-a[j].left))
}

type segsSortByLeftDesc []*segment

func (a segsSortByLeftDesc) Len() int           { return len(a) }
func (a segsSortByLeftDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a segsSortByLeftDesc) Less(i, j int) bool { return a[i].left > a[j].left }

// ========== heleper ==========

// func printSegments(segs []*segment) {
// 	fmt.Print("[")
// 	for i, seg := range segs {
// 		if seg.right == 0 {
// 			continue
// 		}
// 		fmt.Printf("%c:%v,", byte('a'+i), seg)
// 	}
// 	fmt.Println("]")
// }
