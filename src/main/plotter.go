package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"strconv"
	"time"
)

func saveBubbleSortPlot(bestCase plotter.XYs, avgCase plotter.XYs, worstCase plotter.XYs) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Laufzeitverhalten Bubble Sort-Algorithmus"
	p.X.Label.Text = "Problemgröße - Elemente im Array"
	p.Y.Label.Text = "benötigte Zeit - ns"

	err = plotutil.AddLinePoints(p,
		"best case", bestCase,
		"average case", avgCase,
		"worst case", worstCase)
	if err != nil {
		panic(err)
	}

	if err := p.Save(10*vg.Inch, 10*vg.Inch, "bubble.png"); err != nil {
		panic(err)
	}
}

func saveIsPrimePlot(bestCase time.Duration, avgCase time.Duration, worstCase time.Duration, primes isPrimeStruct) {
	durations := plotter.Values{float64(bestCase), float64(avgCase), float64(worstCase)}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Prime-Algorithmus"
	p.X.Label.Text = "Primzahl"
	p.Y.Label.Text = "Laufzeit - ns"

	w := vg.Points(20)

	bars, err := plotter.NewBarChart(durations, w)
	if err != nil {
		panic(err)
	}

	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(0)

	p.Add(bars)
	p.NominalX("best case "+strconv.Itoa(primes.bestCase), "average case "+strconv.Itoa(primes.avgCase), "worst case "+strconv.Itoa(primes.worstCase))

	if err := p.Save(10*vg.Inch, 10*vg.Inch, "isPrime.png"); err != nil {
		panic(err)
	}
}

func saveBinSearchPlot(bestCase time.Duration, avgCase time.Duration, worstCase time.Duration, keys binSearchStruct) {
	durations := plotter.Values{float64(bestCase), float64(avgCase), float64(worstCase)}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Binary Search-Algorithmus"
	p.X.Label.Text = "Key"
	p.Y.Label.Text = "Laufzeit - ns"

	w := vg.Points(20)

	bars, err := plotter.NewBarChart(durations, w)
	if err != nil {
		panic(err)
	}

	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(0)

	p.Add(bars)
	p.NominalX("best case "+strconv.Itoa(keys.bestCaseKey), "average case "+strconv.Itoa(keys.avgCaseKey), "worst case "+strconv.Itoa(keys.worstCaseKey))

	if err := p.Save(4*vg.Inch, 6*vg.Inch, "binSearch.png"); err != nil {
		panic(err)
	}
}

func makePlotable(data []time.Duration) plotter.XYs {
	plotableData := make(plotter.XYs, len(data))
	for i := range data {
		plotableData[i].X = float64(i)
		plotableData[i].Y = float64(data[i])
	}
	return plotableData
}
