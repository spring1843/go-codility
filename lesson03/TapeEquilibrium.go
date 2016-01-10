package lesson03

import (
	"math"
)

// Original Problem
// https://codility.com/programmers/task/tape_equilibrium/
func TapeEquilibrium(A []int) int {
	var mindiff float64 = 1001
	N := len(A)
	P := 0
	headSum := 0
	tailSum := getArraySum(A)
	for P < N-1 {
		headSum += A[P]
		tailSum -= A[P]
		diff := math.Abs(float64(tailSum - headSum))
		if diff < mindiff {
			mindiff = diff
		}
		P++
	}

	return int(mindiff)
}

func getArraySum(A []int) int {
	sum := 0
	i := 0
	for i < len(A) {
		sum += A[i]
		i++
	}
	return sum
}
