package main

import (
	"container/list"
)

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	probMat := make(map[int](map[int]float64))
	for i := 0; i < n; i++ {
		probMat[i] = make(map[int]float64)
	}
	for i := 0; i < len(edges); i++ {
		probMat[edges[i][0]][edges[i][1]] = succProb[i]
		probMat[edges[i][1]][edges[i][0]] = succProb[i]
	}

	queue := list.New()
	inq := make(map[int]bool)
	for ed := range probMat[start] {
		queue.PushBack(ed)
		inq[ed] = true
	}
	for queue.Len() > 0 {
		node := queue.Front()
		queue.Remove(node)
		mid := node.Value.(int)
		delete(inq, mid)
		for ed := range probMat[mid] {
			newProb := probMat[start][mid] * probMat[mid][ed]
			if probMat[start][ed] < newProb {
				probMat[start][ed] = newProb
				if !inq[ed] {
					queue.PushBack(ed)
					inq[ed] = true
				}
			}
		}
	}

	return probMat[start][end]
}
