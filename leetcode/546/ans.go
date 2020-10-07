package main

func removeBoxes(boxes []int) int {
	bb = boxes
	cache = make(map[subProb]int)
	return maxPoints(&subProb{0, len(boxes) - 1, 0})
}

type subProb struct {
	st int
	ed int
	k  int
}

var (
	bb    []int
	cache map[subProb]int
)

func maxPoints(sp *subProb) int {
	if sp.st > sp.ed {
		return 0
	}

	if max, ok := cache[*sp]; ok {
		return max
	}

	for sp.st < sp.ed && bb[sp.ed-1] == bb[sp.ed] {
		sp.ed--
		sp.k++
	}

	max := maxPoints(&subProb{sp.st, sp.ed - 1, 0}) + (sp.k+1)*(sp.k+1)
	for i := sp.st; i < sp.ed; i++ {
		if bb[i] == bb[sp.ed] {
			point := maxPoints(&subProb{sp.st, i, sp.k + 1}) + maxPoints(&subProb{i + 1, sp.ed - 1, 0})
			if max < point {
				max = point
			}
		}
	}
	cache[*sp] = max
	return max
}
