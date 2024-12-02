package day1_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/arnoschutijzer/aoc2024/day1"
	"github.com/stretchr/testify/assert"
)

func TestDay1ThrowsErrorWhenListsAreNotSameLength(t *testing.T) {
	_, err := day1.CalculateTotalDistanceBetween([]int{1}, []int{1, 2})
	assert.ErrorIs(t, err, day1.ErrListsAreNotSameLength)
}

func TestDay1ReturnsTotalDistanceBetween2EqualLists(t *testing.T) {
	distance, err := day1.CalculateTotalDistanceBetween([]int{1, 2, 3, 4}, []int{5, 6, 7, 8})
	assert.Nil(t, err)
	assert.Equal(t, 16, distance)
}

func TestDay1(t *testing.T) {
	firstList, secondList := parseInput()

	distance, err := day1.CalculateTotalDistanceBetween(firstList, secondList)

	assert.Nil(t, err)
	assert.Equal(t, 2769675, distance)
}

func TestDay1ReturnsTotalSimilarityBetween2Lists(t *testing.T) {
	similarity := day1.CalculateTotalSimilarityBetween([]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3})
	assert.Equal(t, 31, similarity)
}

func TestDay1CalculatesSimilarityFromInput(t *testing.T) {
	firstList, secondList := parseInput()

	similarity := day1.CalculateTotalSimilarityBetween(firstList, secondList)
	assert.Equal(t, 24643097, similarity)
}

func parseInput() ([]int, []int) {
	bytes, _ := os.ReadFile("./input.txt")

	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	firstList := make([]int, 0)
	secondList := make([]int, 0)

	for i := range lines {
		line := lines[i]

		columns := strings.Split(line, "   ")
		firstValue, _ := strconv.Atoi(columns[0])
		firstList = append(firstList, firstValue)

		secondValue, _ := strconv.Atoi(columns[1])
		secondList = append(secondList, secondValue)
	}

	return firstList, secondList
}
