package day1

import (
	"errors"
	"sort"
)

var ErrListsAreNotSameLength = errors.New("lists are not the same length")

func CalculateTotalDistanceBetween(firstList []int, secondList []int) (int, error) {
	if len(firstList) != len(secondList) {
		return 0, ErrListsAreNotSameLength
	}

	sort.IntSlice.Sort(firstList)
	sort.IntSlice.Sort(secondList)

	distance := 0
	for k := range firstList {
		distance += absDifference(firstList[k], secondList[k])
	}

	return distance, nil
}

func absDifference(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func CalculateTotalSimilarityBetween(firstList []int, secondList []int) int {
	uniqueIdsFromFirstList := make(map[int]int)

	for _, id := range firstList {
		uniqueIdsFromFirstList[id] = 0
	}

	for _, id := range secondList {
		_, ok := uniqueIdsFromFirstList[id]
		if ok {
			uniqueIdsFromFirstList[id] = uniqueIdsFromFirstList[id] + 1
		}
	}

	similarities := make(map[int]int)
	for key, value := range uniqueIdsFromFirstList {
		similarities[key] += key * value
	}

	similarity := 0
	for _, id := range firstList {
		similarity += similarities[id]
	}

	return similarity
}
