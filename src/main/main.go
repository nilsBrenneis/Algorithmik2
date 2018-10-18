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
		bubbleSort(testData)
		durations[i] = time.Since(start)
	}
	return durations
}

func measure() (durationResults, durationResults, durationResults) {
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

	return resultsBinSearch, resultsIsPrime, resultsBubbleSort
}

func genAverage(durations []time.Duration) (average time.Duration) {
	for i := 0; i < len(durations); i++ {
		average += durations[i]
	}
	tempAvg := float64(average) / float64(len(durations))
	return time.Duration(tempAvg)
}

func main() {
	binSearchDurations, isPrimeDurations, bubbleSortDurations := measure()

	fmt.Println(binSearchString, bestCaseString, genAverage(binSearchDurations.bestCase))
	fmt.Println(binSearchString, avgCaseString, genAverage(binSearchDurations.avgCase))
	fmt.Println(binSearchString, worstCaseString, genAverage(binSearchDurations.worstCase))

	fmt.Println(isPrimeString, bestCaseString, genAverage(isPrimeDurations.bestCase))
	fmt.Println(isPrimeString, avgCaseString, genAverage(isPrimeDurations.avgCase))
	fmt.Println(isPrimeString, worstCaseString, genAverage(isPrimeDurations.worstCase))

	fmt.Println(bubbleSortString, bestCaseString, genAverage(bubbleSortDurations.bestCase))
	fmt.Println(bubbleSortString, avgCaseString, genAverage(bubbleSortDurations.avgCase))
	fmt.Println(bubbleSortString, worstCaseString, genAverage(bubbleSortDurations.worstCase))
}
