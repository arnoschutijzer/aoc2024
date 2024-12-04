package day2

func GetTotalSafeLevels(reports []Report) int {
	safeReports := 0
	for _, report := range reports {

		if report.IsSafe() {
			safeReports += 1
		}
	}

	return safeReports
}

type Report struct {
	Levels []int
}

func (l *Report) IsSafe() bool {
	comparator := getIsSafeComparator(l.Levels[0], l.Levels[1])

	for i := 0; i < len(l.Levels)-1; i++ {
		if !(comparator(l.Levels[i], l.Levels[i+1])) {
			return false
		}
	}

	return true
}

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
