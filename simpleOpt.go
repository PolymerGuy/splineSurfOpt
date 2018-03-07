package main

import (
	"fmt"
	"github.com/PolymerGuy/tinyOpt/plottool"
	"github.com/pkelchte/spline"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot/plotter"
	"math"
	"log"
)

func main() {
	xs := arange(-5, 5, 0.1)
	ys := []float64{}
	for _, x := range xs {
		ys = append(ys, gauss(x))
	}
	serie := plottool.MakeXYs(xs, ys)
	maxY, maxX := findMaxima()

	maxima := plottool.MakeXYs([]float64{maxX}, []float64{maxY})
	plottool.PlotSeries([]plotter.XYs{serie, maxima})
	fmt.Println(findMaxima())
}

func gauss(x float64) float64 {
	a := 1.
	b := 0.0
	c := 0.5
	return a * math.Exp(-math.Pow((x-b), 2)/(2*c*c))
}

func arange(min float64, max float64, step float64) []float64 {
	results := []float64{min}
	val := min
	for val <= max {
		val += step
		results = append(results, val)

	}
	return results
}

func findMaxima() (float64, float64) {
	tol :=1e-6
	initialGuess := []float64{-2.0, -1., 0.7, 2.0}
	initialVals := []float64{}
	for _, x := range initialGuess {
		initialVals = append(initialVals, gauss(x))
	}

	vals := initialVals
	inds := initialGuess
	s := spline.Spline{}

	maxit := 10
	it := 0
	for {
		indices := make([]int, len(inds))

		floats.Argsort(inds, indices)
		vals = sortBy(vals, indices)

		// The spline part
		// Make interpolator in the current range
		s.Set_points(inds, vals, true)
		results := []float64{}
		// Seed the current range
		seed := arange(floats.Min(inds), floats.Max(inds), 0.0001)
		// Evaluate interpolator in current range
		for _, x := range seed {
			results = append(results, s.Operate(x))
		}

		maxInd, _ := argMax(results)

		// Find minimum and its index
		val := gauss(seed[maxInd])

		if !containsElementWithinTol(inds,seed[maxInd],tol){
			vals = append(vals, val)
			inds = append(inds, seed[maxInd])
		}

		fmt.Println(val, seed[maxInd])
		it++
		if it >= maxit {
			return val, seed[maxInd]
		}
	}

}

func argMax(vals []float64) (int, float64) {
	max := vals[0]
	maxInd := 0
	for index, val := range vals {
		if val > max {
			max = val
			maxInd = index
		}
	}
	return maxInd, max
}

func sortBy(vals []float64, indices []int) []float64 {
	if !(len(vals)==len(indices)){
		log.Panic("Values and indices are not of equal length")
	}
	organized := make([]float64, len(vals))
	for i, index := range indices {
		organized[i] = vals[index]
	}
	return organized
}

func containsElementWithinTol(values[]float64,val float64,tol float64) bool{
	for _,value:=range values{
		if math.Abs(value-val)<tol{
			return true
		}
	}
	return false
}