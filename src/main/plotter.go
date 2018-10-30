package main

//https://github.com/gonum/plot/
import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math"
	"time"
)

func saveBubbleSortPlot(bestCase plotter.XYs, avgCase plotter.XYs, worstCase plotter.XYs, log10 bool) {
	var filename string
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Laufzeitverhalten Bubble Sort-Algorithmus"
	if log10 {
		p.X.Label.Text = "Problemgröße - Log10(Elemente im Array)"
		p.Y.Label.Text = "benötigte Zeit - Log10(ns)"
		filename = "bubbleSortLog"
	} else {
		p.X.Label.Text = "Problemgröße - Elemente im Array"
		p.Y.Label.Text = "benötigte Zeit - ns"
		filename = "bubbleSort"
	}

	err = plotutil.AddLinePoints(p,
		"best case", bestCase,
		"average case", avgCase,
		"worst case", worstCase)
	if err != nil {
		panic(err)
	}

	if err := p.Save(10*vg.Inch, 10*vg.Inch, filename+".png"); err != nil {
		panic(err)
	}
}

func saveBinSearchPlot(durations plotter.XYs) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Laufzeitverhalten binary Search-Algorithmus"
	p.X.Label.Text = "gesuchtes Element"
	p.Y.Label.Text = "benötigte Zeit - ns"

	err = plotutil.AddLinePoints(p, durations)
	if err != nil {
		panic(err)
	}

	if err := p.Save(10*vg.Inch, 10*vg.Inch, "binSearch.png"); err != nil {
		panic(err)
	}
}

func saveIsPrimePlot(durations plotter.XYs, log10 bool) {
	var filename string
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Laufzeitverhalten is Prime-Algorithmus"
	if log10 {
		p.X.Label.Text = "Problemgröße - Log10(Größe der Primzahl)"
		p.Y.Label.Text = "benötigte Zeit - Log10(ns)"
		filename = "isPrimeLog"
	} else {
		p.X.Label.Text = "Problemgröße - Größe der Primzahl"
		p.Y.Label.Text = "benötigte Zeit - ns"
		filename = "isPrime"
	}

	err = plotutil.AddLinePoints(p, durations)
	if err != nil {
		panic(err)
	}

	if err := p.Save(10*vg.Inch, 10*vg.Inch, filename+".png"); err != nil {
		panic(err)
	}
}

func makeBubbleSortDataPlotable(durations []time.Duration) plotter.XYs {
	plotableData := make(plotter.XYs, len(durations))
	for i := range durations {
		plotableData[i].X = float64(i)
		plotableData[i].Y = float64(durations[i])
	}
	return plotableData
}

func makeBinSearchDataPlotable(durations []time.Duration, keys []int) plotter.XYs {
	plotableData := make(plotter.XYs, len(durations))

	if len(durations) != len(keys) {
		return plotableData
	}

	for i := range durations {
		plotableData[i].X = float64(keys[i])
		plotableData[i].Y = float64(durations[i])
	}
	return plotableData
}

func transformToLog10(plotData plotter.XYs) plotter.XYs {
	for i := range plotData {
		if plotData[i].X == 0 {
			plotData[i].X = 0
		} else {
			plotData[i].X = math.Log10(plotData[i].X)
		}
		if plotData[i].X == 0 {
			plotData[i].Y = 0
		} else {
			plotData[i].Y = math.Log10(plotData[i].Y)
		}

	}
	return plotData
}
