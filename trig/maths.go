package trig

import (
	"math"
)

const epsilon = 1e-10 // Small value for floating point comparison
const pi = 3.14159265358979323846264338327950288419716939937510582097494459

func DegreeToRadian(deg float64) float64 {
	return deg * (pi / 180.0)
}

func RadianToDegree(rad float64) float64 {
	return rad * (180.0 / pi)
}

func Cos(deg float64) float64 {
	rad := DegreeToRadian(deg)
	return CosRad(rad)
}

func CosRad(rad float64) float64 {
	cosVal := 1.0
	term := 1.0
	for i := 1; Abs(term) > epsilon; i++ {
		term *= -rad * rad / float64((2*i)*(2*i-1))
		cosVal += term
	}
	return cosVal
}

func Abs(term float64) float64 {
	if term > 0 {
		return term
	}
	return term * -1.0
}

func Sin(deg float64) float64 {
	rad := DegreeToRadian(deg)
	return SinRad(rad)
}

func SinRad(rad float64) float64 {
	sinVal := rad
	term := rad
	for i := 1; math.Abs(term) > epsilon; i++ {
		term *= -rad * rad / float64((2*i)*(2*i+1))
		sinVal += term
	}
	return sinVal
}

func Tan(deg float64) float64 {
	rad := DegreeToRadian(deg)
	return TanRad(rad)
}

func TanRad(rad float64) float64 {
	return SinRad(rad) / CosRad(rad)
}

func Asin(val float64) float64 {
	rad := math.Asin(val)
	return RadianToDegree(rad)
}

func Acos(val float64) float64 {
	rad := math.Acos(val)
	return RadianToDegree(rad)
}

func Atan(val float64) float64 {
	rad := math.Atan(val)
	return RadianToDegree(rad)
}
func Pow(number float64, power int) float64 {
	total := 1.0
	for i := 0; i < power; i++ {
		total *= number
	}
	return total
}
func Sqrt(number float64) float64 {
	x := 1.0
	for i := 0; i < 100; i++ {
		x = (x + number/x) / 2.0
	}
	return x
}
