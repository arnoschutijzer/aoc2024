package day2_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/arnoschutijzer/aoc2024/day2"
	"github.com/stretchr/testify/assert"
)

func TestIsSafeForFirstLineInExample(t *testing.T) {
	safe := day2.IsSafe([]int{7, 6, 4, 2, 1})
	assert.True(t, safe)
}

func TestIsUnsafeWhenIncreasingWithFive(t *testing.T) {
	safe := day2.IsSafe([]int{1, 2, 7, 8, 9})
	assert.False(t, safe)
}

func TestIsUnsafeWhenDecreasingWithFour(t *testing.T) {
	safe := day2.IsSafe([]int{9, 7, 6, 2, 1})
	assert.False(t, safe)
}

func TestIsUnsafeWhenFirstIncreasingAndThenDecreasing(t *testing.T) {
	safe := day2.IsSafe([]int{1, 3, 2, 4, 5})
	assert.False(t, safe)
}

func TestIsUnsafeWhenNoIncreaseOrDecrease(t *testing.T) {
	safe := day2.IsSafe([]int{8, 6, 4, 4, 1})
	assert.False(t, safe)
}

func TestFailing(t *testing.T) {
	safe := day2.IsSafe([]int{42, 44, 46, 48, 55})
	assert.False(t, safe)
}

func TestGetNumberOfSafeReports(t *testing.T) {
	report := readReport("./input.txt")
	safeReports := day2.CountSafeReports(report)
	assert.Equal(t, 236, safeReports)
}

func TestGetNumberOfSafeReportsFromExample(t *testing.T) {
	report := readReport("./example.txt")
	safeReports := day2.CountSafeReports(report)
	assert.Equal(t, 2, safeReports)
}

func readReport(path string) [][]int {
	bytes, _ := os.ReadFile(path)
	content := string(bytes)

	lines := strings.Split(content, "\n")

	report := make([][]int, 0)

	for _, line := range lines {
		numbersAsText := strings.Split(line, " ")

		numbersAsInt := make([]int, 0)
		for _, aNumber := range numbersAsText {
			number, _ := strconv.Atoi(aNumber)
			numbersAsInt = append(numbersAsInt, number)
		}

		report = append(report, numbersAsInt)
	}

	return report
}
