package main

import (
	"fmt"
	"math"
)

func levenshteinDistance(firstString, secondString string) int {
	m, n := len(firstString), len(secondString)

	dp := buildDistanceMatrix(m, n)

	for idx := 1; idx < m+1; idx++ {
		for jdx := 1; jdx < n+1; jdx++ {
			cost := 1
			if firstString[idx-1] == secondString[jdx-1] {
				cost = 0
			}
			dp[idx][jdx] = min(
				dp[idx-1][jdx]+1,      // deletion
				dp[idx][jdx-1]+1,      // insertion
				dp[idx-1][jdx-1]+cost, // substitution
			)
		}
	}

	return dp[m][n]
}

func min(args ...int) int {

	_min := math.MaxInt
	for _, num := range args {
		if num < _min {
			_min = num
		}
	}
	return _min
}

func buildDistanceMatrix(m, n int) [][]int {
	dp := make([][]int, m+1)

	for idx := range dp {
		dp[idx] = make([]int, n+1)
	}

	for idx := 0; idx <= m; idx++ {
		dp[idx][0] = idx
	}

	for jdx := 0; jdx <= n; jdx++ {
		dp[0][jdx] = jdx
	}

	return dp
}

func main() {
	firstString := "cooking"
	secondString := "looking"

	distance := levenshteinDistance(firstString, secondString)
	fmt.Printf("The Levenshtein distance between '%s' and '%s' is: %d", firstString, secondString, distance)
}
