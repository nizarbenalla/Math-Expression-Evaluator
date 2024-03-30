package trig

const error = 1.0e-10

func Pow(number float64, power int) float64 {
	total := 1.0
	for i := 0; i < power; i++ {
		total *= number
	}
	return total
}
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(number float64) float64 {
	x := 1.0
	for i := 0; i < 100; i++ {
		x = (x + number/x) / 2.0
	}
	return x
}

func Cos(x float64) float64 {
	result := 1.0
	t := 1.0
	for i := 1; i <= 10; i++ {
		t *= (-1) * x * x / float64(2*i) / float64(2*i-1)
		result += t
	}
	return result
}

func Acos(x float64) float64 {
	xn := x
	xn1 := x
	for {
		f := Cos(xn) - x
		fp := -Sin(xn)
		xn1 = xn - f/fp
		if abs(xn1-xn) < error {
			break
		}
		xn = xn1
	}
	return xn1
}

func Sin(x float64) float64 {
	result := x
	t := x
	for i := 1; i <= 10; i++ {
		t *= (-1) * x * x / float64(2*i) / float64(2*i+1)
		result += t
	}
	return result
}

func Asin(x float64) float64 {
	xn := x
	xn1 := x
	for {
		f := Sin(xn) - x
		fp := Cos(xn)
		xn1 = xn - f/fp
		if abs(xn1-xn) < error {
			break
		}
		xn = xn1
	}
	return xn1
}

func Tan(x float64) float64 {
	return Sin(x) / Cos(x)
}

func Atan(x float64) float64 {
	xn := x
	xn1 := x
	for {
		f := Tan(xn) - x
		fp := 1 / (Cos(xn) * Cos(xn))
		xn1 = xn - f/fp
		if abs(xn1-xn) < error {
			break
		}
		xn = xn1
	}
	return xn1
}
