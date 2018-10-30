package main

import (
	"time"
)

var amountMeasurements = 100

//runs multiple times with same data input to generate average runtime
func measureBinSearchRun(testData []int, key int) time.Duration {
	durations := make([]time.Duration, amountMeasurements)

	for i := 0; i < len(durations); i++ {
		start := time.Now()
		binarySearch(testData, key)
		durations[i] = time.Since(start)
	}
	return genAverage(durations)
}

//runs multiple times with same data input to generate average runtime
func measureBubbleSortRun(testData []int) time.Duration {
	durations := make([]time.Duration, amountMeasurements)
	for i := 0; i < len(durations); i++ {
		start := time.Now()
		bubbleSort(testData)
		durations[i] = time.Since(start)
	}
	return genAverage(durations)
}

//runs multiple times with same data input to generate average runtime
func measureIsPrimeRun(p int) time.Duration {
	durations := make([]time.Duration, amountMeasurements)
	for i := 0; i < len(durations); i++ {
		start := time.Now()
		isPrime(p)
		durations[i] = time.Since(start)
	}
	return genAverage(durations)
}

//run binary search with every key in array
func measureBinSearch(testData binSearchStruct) []time.Duration {
	avgDurations := make([]time.Duration, len(testData.keys))
	for i := 0; i < len(testData.keys); i++ {
		avgDurations[i] = measureBinSearchRun(testData.arr, testData.keys[i])
	}
	return avgDurations
}

//increases integer to check if it is prime
func measureIsPrime() []time.Duration {
	maxPotentialPrime := 2000
	avgDurations := make([]time.Duration, maxPotentialPrime)
	for i := 0; i < maxPotentialPrime; i++ {
		avgDurations[i] = measureIsPrimeRun(i)
	}
	return avgDurations
}

//increase slice length
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
	average := time.Duration(float64(accu) / float64(len(durations)))
	//get rid of outliers greater than value 'percentage'
	accu = 0
	countDurations := 0
	percentage := 0.75
	threshold := time.Duration(float64(average) * percentage)
	for i := 0; i < len(durations); i++ {
		if durations[i]-average < threshold && average-durations[i] < threshold {
			accu += durations[i]
			countDurations++
		}
	}
	return time.Duration(float64(accu) / float64(countDurations))
}

func createBubbleSortPlot() {
	testData := genDataBubbleSort()

	durationsBestCase := measureBubbleSort(testData.bestCase)
	durationsAvgCase := measureBubbleSort(testData.avgCase)
	durationsWorstCase := measureBubbleSort(testData.worstCase)

	bestCasePlot := makeBubbleSortDataPlotable(durationsBestCase)
	avgCasePlot := makeBubbleSortDataPlotable(durationsAvgCase)
	worstCasePlot := makeBubbleSortDataPlotable(durationsWorstCase)

	saveBubbleSortPlot(bestCasePlot, avgCasePlot, worstCasePlot, false)

	bestCaseLogPlot := transformToLog10(bestCasePlot)
	avgCaseLogPlot := transformToLog10(avgCasePlot)
	worstCaseLogPlot := transformToLog10(worstCasePlot)

	saveBubbleSortPlot(bestCaseLogPlot, avgCaseLogPlot, worstCaseLogPlot, true)
}

func createIsPrimePlot() {
	//no need for test data creation. measureIsPrime() gets test data from iteration
	durations := measureIsPrime()

	plot := makeBubbleSortDataPlotable(durations)
	saveIsPrimePlot(plot, false)

	logPlot := transformToLog10(plot)
	saveIsPrimePlot(logPlot, true)
}

func createBinSearchPlot() {
	testData := genDataBinSearch()
	durations := measureBinSearch(testData)
	saveBinSearchPlot(makeBinSearchDataPlotable(durations, testData.keys))
}

func main() {
	createBubbleSortPlot()
	createIsPrimePlot()
	createBinSearchPlot()
}
