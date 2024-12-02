package day1_test

import (
	"fmt"
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
	bytes, err := os.ReadFile("./input.txt")
	assert.Nil(t, err)

	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	firstList := make([]int, 0)
	secondList := make([]int, 0)

	for i := range lines {
		line := lines[i]

		fmt.Println(line)

		// skip empty lines
		if len(line) == 0 {
			continue
		}

		columns := strings.Split(line, "   ")
		firstValue, err := strconv.Atoi(columns[0])
		assert.Nil(t, err)
		firstList = append(firstList, firstValue)

		secondValue, err := strconv.Atoi(columns[1])
		assert.Nil(t, err)
		secondList = append(secondList, secondValue)
	}

	distance, err := day1.CalculateTotalDistanceBetween(firstList, secondList)

	assert.Nil(t, err)
	assert.Equal(t, 2769675, distance)
}
