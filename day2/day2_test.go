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
	report := day2.Report{Levels: []int{7, 6, 4, 2, 1}}
	assert.True(t, report.IsSafe())
}

func TestIsUnsafeWhenIncreasingWithFive(t *testing.T) {
	report := day2.Report{Levels: []int{1, 2, 7, 8, 9}}
	assert.False(t, report.IsSafe())
}

func TestIsUnsafeWhenDecreasingWithFour(t *testing.T) {
	report := day2.Report{Levels: []int{9, 7, 6, 2, 1}}
	assert.False(t, report.IsSafe())
}

func TestIsUnsafeWhenFirstIncreasingAndThenDecreasing(t *testing.T) {
	report := day2.Report{Levels: []int{1, 3, 2, 4, 5}}
	assert.False(t, report.IsSafe())
}

func TestIsUnsafeWhenNoIncreaseOrDecrease(t *testing.T) {
	report := day2.Report{Levels: []int{8, 6, 4, 4, 1}}
	assert.False(t, report.IsSafe())
}

func TestFailing(t *testing.T) {
	report := day2.Report{Levels: []int{42, 44, 46, 48, 55}}
	assert.False(t, report.IsSafe())
}

func TestGetNumberOfSafeReports(t *testing.T) {
	reports := readReport("./input.txt")
	safeReports := day2.GetTotalSafeLevels(reports)
	assert.Equal(t, 236, safeReports)
}

func TestGetNumberOfSafeReportsFromExample(t *testing.T) {
	reports := readReport("./example.txt")
	safeReports := day2.GetTotalSafeLevels(reports)
	assert.Equal(t, 2, safeReports)
}

func readReport(path string) []day2.Report {
	bytes, _ := os.ReadFile(path)
	content := string(bytes)

	lines := strings.Split(content, "\n")

	reports := []day2.Report{}
	for _, line := range lines {
		numbersAsText := strings.Split(line, " ")

		numbersAsInt := make([]int, 0)
		for _, aNumber := range numbersAsText {
			number, _ := strconv.Atoi(aNumber)
			numbersAsInt = append(numbersAsInt, number)
		}

		reports = append(reports, day2.Report{Levels: numbersAsInt})
	}

	return reports
}
