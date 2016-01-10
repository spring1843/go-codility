package lesson03

// Original Problem
// https://codility.com/programmers/task/perm_missing_elem/
func PermMissingElem(N int, A []int) []int {
	counters := make([]int, N)

	maxCounter := 0
	i := 0
	ii := 0
	M := len(A) - 1
	for i <= M {
		if A[i] == N+1 {
			ii = 0
			for ii < N {
				counters[ii] = maxCounter
				ii++
			}
		}
		if A[i] >= 1 && A[i] <= N {
			counters[A[i]-1]++
			if counters[A[i]-1] > maxCounter {
				maxCounter = counters[A[i]-1]
			}
		}
		i++
	}

	return counters
}
