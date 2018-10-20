package main

import (
	"fmt"
	"time"
)

var amountMeasurements = 100

func measureBinSearchRun(testData []int, key int) time.Duration {
	durations := make([]time.Duration, amountMeasurements)

	for i := 0; i < len(durations); i++ {
		start := time.Now()
		binarySearch(testData, key)
		durations[i] = time.Since(start)
	}
	return genAverage(durations)
}

func measureIsPrimeRun(p int) time.Duration {
	durations := make([]time.Duration, amountMeasurements)
	for i := 0; i < len(durations); i++ {
		start := time.Now()
		isPrime(p)
		durations[i] = time.Since(start)
	}
	return genAverage(durations)
}

func measureBubbleSortRun(testData []int) time.Duration {
	durations := make([]time.Duration, amountMeasurements)
	for i := 0; i < len(durations); i++ {
		start := time.Now()
		bubbleSort(testData)
		durations[i] = time.Since(start)
	}
	return genAverage(durations)
}

func measureBubbleSort(testData []int) []time.Duration {
	avgDurations := make([]time.Duration, len(testData))
	for i := 0; i < len(testData); i++ {
		avgDurations[i] = measureBubbleSortRun(testData[:i])
	}
	return avgDurations
}

func genAverage(durations []time.Duration) time.Duration {
	var accu time.Duration
	for i := 0; i < len(durations); i++ {
		accu += durations[i]
	}
	return time.Duration(float64(accu) / float64(len(durations)))
}

func main() {
	bubbleSortTestData := genDataBubbleSort()

	durationsBestCaseBubbleSort := measureBubbleSort(bubbleSortTestData.bestCase)
	durationsAvgCaseubbleSort := measureBubbleSort(bubbleSortTestData.avgCase)
	durationsWorstCaseubbleSort := measureBubbleSort(bubbleSortTestData.worstCase)
	fmt.Println(durationsBestCaseBubbleSort)
	fmt.Println(durationsAvgCaseubbleSort)
	fmt.Println(durationsWorstCaseubbleSort)

	isPrimeTestData := genDataIsPrime()

	durationsBestCaseIsPrime := measureIsPrimeRun(isPrimeTestData.bestCase)
	durationsAvgCaseIsPrime := measureIsPrimeRun(isPrimeTestData.avgCase)
	durationsWorstCaseIsPrime := measureIsPrimeRun(isPrimeTestData.worstCase)
	fmt.Println(durationsBestCaseIsPrime)
	fmt.Println(durationsAvgCaseIsPrime)
	fmt.Println(durationsWorstCaseIsPrime)

	binSearchTestData := genDataBinSearch()

	durationsBestCaseBinSearch := measureBinSearchRun(binSearchTestData.arr, binSearchTestData.bestCaseKey)
	durationsAvgCaseBinSearch := measureBinSearchRun(binSearchTestData.arr, binSearchTestData.avgCaseKey)
	durationsWorstCaseBinSearch := measureBinSearchRun(binSearchTestData.arr, binSearchTestData.worstCaseKey)
	fmt.Println(durationsBestCaseBinSearch)
	fmt.Println(durationsAvgCaseBinSearch)
	fmt.Println(durationsWorstCaseBinSearch)
}
