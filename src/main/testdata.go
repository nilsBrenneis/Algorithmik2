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
	isPrimeTestData.worstCase = 21384943
}

func genDataBubbleSort() {
	bubbleSortTestData.bestCase = []int{1, 2, 89, 100, 180, 181, 1000, 1045, 1056, 1100, 3000, 3001, 3453, 3679, 4000}

	maxInt := 100
	bubbleSortTestData.avgCase = []int{rand.Intn(maxInt), rand.Intn(maxInt), rand.Intn(maxInt), rand.Intn(maxInt),
		rand.Intn(maxInt), rand.Intn(maxInt), rand.Intn(maxInt), rand.Intn(maxInt), rand.Intn(maxInt), rand.Intn(maxInt),
		rand.Intn(maxInt), rand.Intn(maxInt), rand.Intn(maxInt), rand.Intn(maxInt), rand.Intn(maxInt)}

	bubbleSortTestData.worstCase = []int{4000, 3679, 3453, 3001, 3000, 1100, 1056, 1045, 1000, 181, 180, 100, 89, 2, 1}
}
