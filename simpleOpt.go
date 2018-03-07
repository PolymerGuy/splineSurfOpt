package main

import (
	"fmt"
	"github.com/PolymerGuy/tinyOpt/plottool"
	"github.com/pkelchte/spline"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot/plotter"
	"github.com/PolymerGuy/splineSurfOpt/functions"
	"github.com/PolymerGuy/splineSurfOpt/maths"
)

func main() {
	function := functions.Himmelbau
	xs := maths.Arange(-5, 5., 0.1)
	ys := []float64{}
	for _, x := range xs {
		ys = append(ys, function(x))
	}
	serie := plottool.MakeXYs(xs, ys)

	initialGuess := maths.Arange(-5., 5., 2)
	maxY, maxX := minimize(function,initialGuess)

	maxima := plottool.MakeXYs([]float64{maxX}, []float64{maxY})
	plottool.PlotSeries([]plotter.XYs{serie, maxima},"Results")

	fmt.Println(maths.ArgMin(ys))
}



func minimize(f func(float64)float64,initialGuess []float64) (float64, float64) {
	tol :=1e-6

	fmt.Println(initialGuess)
	initialVals := []float64{}
	for _, x := range initialGuess {
		initialVals = append(initialVals, f(x))
	}

	vals := initialVals
	inds := initialGuess
	s := spline.Spline{}

	maxit := 8
	it := 0
	for {
		indices := make([]int, len(inds))

		floats.Argsort(inds, indices)
		vals = maths.SortBy(vals, indices)

		// The spline part
		// Make interpolator in the current range
		s.Set_points(inds, vals, true)
		results := []float64{}
		// Seed the current range
		seed := maths.Arange(floats.Min(inds), floats.Max(inds), 0.0001)
		// Evaluate interpolator in current range
		for _, x := range seed {
			results = append(results, s.Operate(x))
		}

		maxInd, _ := maths.ArgMin(results)

		// Find minimum and its index
		val := f(seed[maxInd])

		if !maths.ContainsElementWithinTol(inds,seed[maxInd],tol){
			vals = append(vals, val)
			inds = append(inds, seed[maxInd])
		}

		fmt.Println(val, seed[maxInd])
		//serie := plottool.MakeXYs(seed,results)
		//filename := "Results"+fmt.Sprint(it)
		//plottool.PlotSeries([]plotter.XYs{serie},filename)
		it++
		if it >= maxit {
			return val, seed[maxInd]
		}
	}

}


