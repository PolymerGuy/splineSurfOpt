package functions

import "math"

func Gauss(x float64) float64 {
	a := 1.
	b := 0.0
	c := 0.5
	return a * math.Exp(-math.Pow((x-b), 2)/(2*c*c))
}

func SixHumpCamelFunction(x float64) float64 {
	return -(4.0 - 2.1*x*x + math.Pow(x, 4)/3.0) * x * x
}

func Forrester(x float64) float64 {
	return -math.Pow(6.0*x-2.0, 2) * math.Sin(12*x-4)
}

func Himmelbau(x float64) float64 {
	y := 2.0
	return -math.Pow(math.Pow(x, 2)+y-11, 2) - math.Pow(x+math.Pow(y, 2)-7, 2)
}
