package main

import (
	"math/rand"
)

type binSearchStruct struct {
	keys []int
	arr  []int
}

type bubbleSortStruct struct {
	bestCase  []int
	avgCase   []int
	worstCase []int
}

func genDataBinSearch() binSearchStruct {
	arrLen := 100000000
	key := arrLen
	var binSearchTestData binSearchStruct
	binSearchTestData.arr = make([]int, arrLen)
	binSearchTestData.keys = make([]int, 27)

	for i := 0; i < len(binSearchTestData.arr); i++ {
		binSearchTestData.arr[i] = i
	}

	for i := 0; i < len(binSearchTestData.keys); i++ {
		key = key / 2
		binSearchTestData.keys[i] = key
	}
	return binSearchTestData
}

func genDataBubbleSort() bubbleSortStruct {
	arrLen := 1000
	maxRandInt := arrLen * 2
	var bubbleSortTestData bubbleSortStruct
	bubbleSortTestData.bestCase = make([]int, arrLen)
	bubbleSortTestData.avgCase = make([]int, arrLen)
	bubbleSortTestData.worstCase = make([]int, arrLen)

	for i := 1; i <= len(bubbleSortTestData.bestCase); i++ {
		bubbleSortTestData.bestCase[i-1] = i
	}

	for i := 0; i < len(bubbleSortTestData.avgCase); i++ {
		bubbleSortTestData.avgCase[i] = rand.Intn(maxRandInt)
	}

	for i := 0; i < len(bubbleSortTestData.worstCase); i++ {
		bubbleSortTestData.worstCase[i] = len(bubbleSortTestData.worstCase) - i
	}
	return bubbleSortTestData
}
