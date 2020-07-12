package main

import "math"

const (
	precision = 0.00001
)

var (
	// 0,0 starts from left-bottom
	directions = [][]float64{
		{0.0, 1.0},  // right
		{1.0, 0.0},  // up
		{0.0, -1.0}, // left
		{-1.0, 0.0}, // down
	}
)

func getMinDistSum(positions [][]int) float64 {
	n := len(positions)
	if n <= 1 {
		return 0.0
	}
	center := make([]float64, 2)

	// try Manhattan center
	for i := n - 1; i >= 0; i-- {
		center[0] += float64(positions[i][0])
		center[1] += float64(positions[i][1])
	}
	center[0] /= float64(n)
	center[1] /= float64(n)
	minSum := getDistSum(positions, center)

	// try all pos as center
	tryPos := make([]float64, 2)
	for i := n - 1; i >= 0; i-- {
		tryPos[0] = float64(positions[i][0])
		tryPos[1] = float64(positions[i][1])
		sum := getDistSum(positions, tryPos)
		if minSum > sum {
			minSum = sum
			center, tryPos = tryPos, center
		}
	}

	tryStep := 50.0
	for {
		improved := false
		maxDiff := 0.0
		for i := 0; i < 4; i++ {
			tryPos[0] = center[0] + directions[i][0]*tryStep
			tryPos[1] = center[1] + directions[i][1]*tryStep
			sum := getDistSum(positions, tryPos)
			if minSum > sum {
				diff := minSum - sum
				if maxDiff < diff {
					maxDiff = diff
				}
				minSum = sum
				center, tryPos = tryPos, center
				improved = true
				break
			}
			diff := sum - minSum
			if maxDiff < diff {
				maxDiff = diff
			}
		}
		if !improved {
			if maxDiff < precision {
				break
			}
			tryStep /= 2
		}
	}

	return minSum
}

func getDistSum(positions [][]int, center []float64) float64 {
	sum := 0.0
	for i := len(positions) - 1; i >= 0; i-- {
		xd := float64(positions[i][0]) - center[0]
		yd := float64(positions[i][1]) - center[1]
		sum += math.Sqrt(xd*xd + yd*yd)
	}
	return sum
}

// another exit condition
func getMinDistSum2(positions [][]int) float64 {
	n := len(positions)
	if n <= 1 {
		return 0.0
	}
	center := make([]float64, 2)

	// try Manhattan center
	for i := n - 1; i >= 0; i-- {
		center[0] += float64(positions[i][0])
		center[1] += float64(positions[i][1])
	}
	center[0] /= float64(n)
	center[1] /= float64(n)
	minSum := getDistSum(positions, center)

	// try all pos as center
	tryPos := make([]float64, 2)
	for i := n - 1; i >= 0; i-- {
		tryPos[0] = float64(positions[i][0])
		tryPos[1] = float64(positions[i][1])
		sum := getDistSum(positions, tryPos)
		if minSum > sum {
			minSum = sum
			center, tryPos = tryPos, center
		}
	}

	tryStep := 50.0
	for tryStep >= precision {
		improved := false
		for i := 0; i < 4; i++ {
			tryPos[0] = center[0] + directions[i][0]*tryStep
			tryPos[1] = center[1] + directions[i][1]*tryStep
			sum := getDistSum(positions, tryPos)
			if minSum > sum {
				minSum = sum
				center, tryPos = tryPos, center
				improved = true
				break
			}
		}
		if !improved {
			tryStep /= 2
		}
	}

	return minSum
}
