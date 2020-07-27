package main

// 64ms
// 6.2MB
func getLengthOfOptimalCompression(s string, k int) int {
	letters := []byte{'#'}
	counts := []int{0}
	gNum := 0
	// grouping
	for i := range s {
		if s[i] != letters[gNum] {
			// new group
			letters = append(letters, s[i])
			counts = append(counts, 1)
			gNum++
			continue
		}
		counts[gNum]++
	}
	gNum++

	sums := make([]int, gNum)
	for i := 1; i < gNum; i++ {
		sums[i] = sums[i-1] + counts[i]
	}

	// dp[i][cost] means the minLen of group[1] to group[i] with cost
	dp := make([](map[int]int), gNum)
	for i := 0; i < gNum; i++ {
		dp[i] = make(map[int]int)
	}
	lastLetterIdx := make([][26]int, gNum)
	// group '#', cost 0, len 0
	dp[0][0] = 0

	for i := 1; i < gNum; i++ {
		char := letters[i] - 'a'
		count := counts[i]
		mergedGroup := i
		midCost := 0
		for { // mergedGroup
			iLen := 0
			for { // iLen
				currCost := count - ceilings[iLen]
				if currCost < 0 {
					currCost = 0
				}

				for preCost, preLen := range dp[mergedGroup-1] {
					cost := preCost + midCost + currCost
					if cost > k {
						continue
					}
					length := preLen + iLen
					if minLen, ok := dp[i][cost]; !ok || minLen > length {
						dp[i][cost] = length
					}
				}

				if currCost == 0 {
					break
				}
				iLen++
			}
			// merge previous group with same letter from right to left
			mg := mergedGroup
			mergedGroup = lastLetterIdx[mergedGroup-1][char]
			if mergedGroup <= 0 {
				break
			}
			count += counts[mergedGroup]
			// midCost increased by the total length between 2 merged group
			midCost += (sums[mg-1] - sums[mergedGroup])
			if midCost > k {
				break
			}
		}
		copy(lastLetterIdx[i][:], lastLetterIdx[i-1][:])
		lastLetterIdx[i][char] = i
	} // for groups

	minLen := 101
	for _, length := range dp[gNum-1] {
		if minLen > length {
			minLen = length
		}
	}
	return minLen
}

var (
	// ceilings[i] means the up-bound of compressedLen i
	ceilings = [...]int{
		0,
		1,
		9,
		99,
		999,
	}
)
