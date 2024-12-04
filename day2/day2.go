package day2

import "fmt"

func getIsSafeComparator(a int, b int) func(a int, b int) bool {
	isGradually := func(a int, b int) bool {
		return absDifference(a, b) > 0 && absDifference(a, b) <= 3
	}

	if a > b {
		return func(a int, b int) bool {
			return a > b && isGradually(a, b)
		}
	}

	return func(a int, b int) bool {
		return a < b && isGradually(a, b)
	}
}

func absDifference(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func IsSafe(level []int) bool {
	comparator := getIsSafeComparator(level[0], level[1])

	for i := 0; i < len(level)-1; i++ {
		if !(comparator(level[i], level[i+1])) {
			return false
		}
	}

	return true
}

func CountSafeReports(report [][]int) int {
	safeReports := 0
	for index, level := range report {

		if IsSafe(level) {
			fmt.Println(index, level)
			safeReports += 1
		}
	}

	return safeReports
}
