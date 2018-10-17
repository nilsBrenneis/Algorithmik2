package main

import (
	"fmt"
	"time"
)

type durationResults struct {
	bestCase  []time.Duration
	avgCase   []time.Duration
	worstCase []time.Duration
}

var worstCaseString = "worst case:  "
var avgCaseString = "average case:"
var bestCaseString = "best case:   "

var binSearchString = "binary search |"
var isPrimeString = "is prime      |"
var bubbleSortString = "bubble sort   |"

var amountMeasurements = 100

func binSearchAnalysis(testData []int, key int) []time.Duration {
	durations := make([]time.Duration, amountMeasurements)

	for i := 0; i < len(durations); i++ {
		start := time.Now()
		binarySearch(testData, key)
		durations[i] = time.Since(start)
	}
	return durations
}

func isPrimeAnalysis(p int) []time.Duration {
	durations := make([]time.Duration, amountMeasurements)

	for i := 0; i < len(durations); i++ {
		start := time.Now()
		isPrime(p)
		durations[i] = time.Since(start)
	}
	return durations
}

func bubbleSortAnalysis(testData []int) []time.Duration {
	durations := make([]time.Duration, amountMeasurements)

	for i := 0; i < len(durations); i++ {
		start := time.Now()
		bubblesort(testData)
		durations[i] = time.Since(start)
	}
	return durations
}

func main() {
	genDataBinSearch()
	genDataIsPrime()
	genDataBubbleSort()

	bestResBinSearch := binSearchAnalysis(binSearchTestData.arr, binSearchTestData.bestCaseKey)
	avgResBinSearch := binSearchAnalysis(binSearchTestData.arr, binSearchTestData.avgCaseKey)
	worstResBinsearch := binSearchAnalysis(binSearchTestData.arr, binSearchTestData.worstCaseKey)
	resultsBinSearch := durationResults{bestResBinSearch, avgResBinSearch, worstResBinsearch}

	bestResIsPrime := isPrimeAnalysis(isPrimeTestData.bestCase)
	avgResIsPrime := isPrimeAnalysis(isPrimeTestData.avgCase)
	worstResIsPrime := isPrimeAnalysis(isPrimeTestData.worstCase)
	resultsIsPrime := durationResults{bestResIsPrime, avgResIsPrime, worstResIsPrime}

	bestResBubbleSort := bubbleSortAnalysis(bubbleSortTestData.bestCase)
	avgResBubbleSort := bubbleSortAnalysis(bubbleSortTestData.avgCase)
	worstResBubbleSort := bubbleSortAnalysis(bubbleSortTestData.worstCase)
	resultsBubbleSort := durationResults{bestResBubbleSort, avgResBubbleSort, worstResBubbleSort}

	fmt.Println(binSearchString, bestCaseString, resultsBinSearch.bestCase)
	fmt.Println(binSearchString, avgCaseString, resultsBinSearch.avgCase)
	fmt.Println(binSearchString, worstCaseString, resultsBinSearch.worstCase)

	fmt.Println(isPrimeString, bestCaseString, resultsIsPrime.bestCase)
	fmt.Println(isPrimeString, avgCaseString, resultsIsPrime.avgCase)
	fmt.Println(isPrimeString, worstCaseString, resultsIsPrime.worstCase)

	fmt.Println(bubbleSortString, bestCaseString, resultsBubbleSort.bestCase)
	fmt.Println(bubbleSortString, avgCaseString, resultsBubbleSort.avgCase)
	fmt.Println(bubbleSortString, worstCaseString, resultsBubbleSort.worstCase)
}
