package main

import (
	"math/rand"
)

type binSearchStruct struct {
	bestCaseKey  int
	avgCaseKey   int
	worstCaseKey int
	arr          []int
}

type isPrimeStruct struct {
	bestCase  int
	avgCase   int
	worstCase int
}

type bubbleSortStruct struct {
	bestCase  []int
	avgCase   []int
	worstCase []int
}

var binSearchTestData binSearchStruct

var isPrimeTestData isPrimeStruct

func genDataBinSearch() binSearchStruct {
	arrLen := 1000
	binSearchTestData.arr = make([]int, arrLen)

	for i := 0; i < len(binSearchTestData.arr); i++ {
		binSearchTestData.arr[i] = i
	}

	binSearchTestData.bestCaseKey = 500
	binSearchTestData.avgCaseKey = 999
	binSearchTestData.worstCaseKey = 250
	return binSearchTestData
}

func genDataIsPrime() isPrimeStruct {
	isPrimeTestData.bestCase = 1
	isPrimeTestData.avgCase = 65539
	isPrimeTestData.worstCase = 213849
	return isPrimeTestData
}

func genDataBubbleSort() bubbleSortStruct {
	arrLen := 1000
	maxRandInt := 2000
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
