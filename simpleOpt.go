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
	function := himmelbau
	xs := arange(-5, 5., 0.1)
	ys := []float64{}
	for _, x := range xs {
		ys = append(ys, function(x))
	}
	serie := plottool.MakeXYs(xs, ys)

	initialGuess := arange(-5., 5., 2)
	maxY, maxX := findMaxima(function,initialGuess)

	maxima := plottool.MakeXYs([]float64{maxX}, []float64{maxY})
	plottool.PlotSeries([]plotter.XYs{serie, maxima},"Results")

	fmt.Println(argMax(ys))
}

func gauss(x float64) float64 {
	a := 1.
	b := 0.0
	c := 0.5
	return a * math.Exp(-math.Pow((x-b), 2)/(2*c*c))
}

func SixHumpCamelFunction(x float64) float64 {
	return -(4.0-2.1*x*x+math.Pow(x,4)/3.0)*x*x
}


func Forrester(x float64) float64 {
	return -math.Pow(6.0*x-2.0,2)*math.Sin(12*x-4)
}

func arange(min float64, max float64, step float64) []float64 {
	results := []float64{}
	val := min
	for val <= max {
		results = append(results, val)
		val += step


	}
	return results
}

func findMaxima(f func(float64)float64,initialGuess []float64) (float64, float64) {
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
		val := f(seed[maxInd])

		if !containsElementWithinTol(inds,seed[maxInd],tol){
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

func himmelbau(x float64)float64{
	y:=2.0
	return -math.Pow(math.Pow(x,2) +y-11,2)-math.Pow(x+math.Pow(y,2)-7,2)
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