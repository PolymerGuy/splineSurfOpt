package maths

import (
	"log"
	"math"
)

func ArgMax(vals []float64) (int, float64) {
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

func SortBy(vals []float64, indices []int) []float64 {
	if !(len(vals)==len(indices)){
		log.Panic("Values and indices are not of equal length")
	}
	organized := make([]float64, len(vals))
	for i, index := range indices {
		organized[i] = vals[index]
	}
	return organized
}

func ContainsElementWithinTol(values[]float64,val float64,tol float64) bool{
	for _,value:=range values{
		if math.Abs(value-val)<tol{
			return true
		}
	}
	return false
}

func Arange(min float64, max float64, step float64) []float64 {
	results := []float64{}
	val := min
	for val <= max {
		results = append(results, val)
		val += step


	}
	return results
}
