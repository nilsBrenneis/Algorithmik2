package main

import (
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

func createBubbleSortPlot() {
	testData := genDataBubbleSort()

	durationsBestCase := measureBubbleSort(testData.bestCase)
	durationsAvgCase := measureBubbleSort(testData.avgCase)
	durationsWorstCase := measureBubbleSort(testData.worstCase)

	saveBubbleSortPlot(makePlotable(durationsBestCase), makePlotable(durationsAvgCase),
		makePlotable(durationsWorstCase))
}

func createIsPrimePlot() {
	testData := genDataIsPrime()

	durationsBestCase := measureIsPrimeRun(testData.bestCase)
	durationsAvgCase := measureIsPrimeRun(testData.avgCase)
	durationsWorstCase := measureIsPrimeRun(testData.worstCase)

	saveIsPrimePlot(durationsBestCase, durationsAvgCase, durationsWorstCase, testData)
}

func createBinSearchPlot() {
	testData := genDataBinSearch()

	durationsBestCase := measureBinSearchRun(testData.arr, testData.bestCaseKey)
	durationsAvgCase := measureBinSearchRun(testData.arr, testData.avgCaseKey)
	durationsWorstCase := measureBinSearchRun(testData.arr, testData.worstCaseKey)

	saveBinSearchPlot(durationsBestCase, durationsAvgCase, durationsWorstCase, testData)
}

//https://github.com/gonum/plot/
func main() {
	createBubbleSortPlot()
	createIsPrimePlot()
	createBinSearchPlot()
}
