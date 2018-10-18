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

var bubbleSortTestData bubbleSortStruct

func genDataBinSearch() {
	binSearchTestData.arr = []int{1, 2, 89, 100, 180, 181, 1000, 1045, 1056, 1100, 3000, 3001, 3453, 3679, 4000}
	binSearchTestData.bestCaseKey = 1045
	binSearchTestData.avgCaseKey = 100
	binSearchTestData.worstCaseKey = 1
}

func genDataIsPrime() {
	isPrimeTestData.bestCase = 1
	isPrimeTestData.avgCase = 65539
	isPrimeTestData.worstCase = 213849
}

func genDataBubbleSort() {
	arrLen := 1000
	bubbleSortTestData.bestCase = make([]int, arrLen)
	for i := 1; i <= len(bubbleSortTestData.bestCase); i++ {
		bubbleSortTestData.bestCase[i-1] = i
	}

	maxRandInt := 2000
	bubbleSortTestData.avgCase = make([]int, arrLen)
	for i := 0; i < len(bubbleSortTestData.avgCase); i++ {
		bubbleSortTestData.avgCase[i] = rand.Intn(maxRandInt)
	}

	bubbleSortTestData.worstCase = make([]int, arrLen)
	for i := 0; i < len(bubbleSortTestData.worstCase); i++ {
		bubbleSortTestData.worstCase[i] = len(bubbleSortTestData.worstCase) - i
	}
}
