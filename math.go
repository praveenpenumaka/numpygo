package numpygo

import (
	"math"
)

// EACH ELEMENT - lambda functions
func ForEachElement(nd NDArray, lambda func(v float64) float64) NDArray {
	newArr := newNDArray(nd.DType, nd.Shape)
	for index, value := range nd.Elements.Values {
		newArr.Elements.Values[index] = lambda(value)
	}
	return newArr
}

func Exp(nd NDArray) NDArray {
	return ForEachElement(nd, func(v float64) float64 {
		return math.Exp(v)
	})
}

func Log2(nd NDArray) NDArray {
	return ForEachElement(nd, func(v float64) float64 {
		return math.Log(v) / math.Log(2)
	})
}

func Log(nd NDArray) NDArray {
	return ForEachElement(nd, func(v float64) float64 {
		return math.Log(v)
	})
}

func Add(nd NDArray, factor float64) NDArray {
	return ForEachElement(nd, func(v float64) float64 {
		return v + factor
	})
}

func Sub(nd NDArray, factor float64) NDArray {
	return ForEachElement(nd, func(v float64) float64 {
		return v - factor
	})
}

func Mult(nd NDArray, factor float64) NDArray {
	return ForEachElement(nd, func(v float64) float64 {
		return v * factor
	})
}

func Div(nd NDArray, factor float64) NDArray {
	return ForEachElement(nd, func(v float64) float64 {
		return v / factor
	})
}

func Pow(nd NDArray, factor float64) NDArray {
	return ForEachElement(nd, func(v float64) float64 {
		return math.Pow(v, factor)
	})
}

func Equal(nd NDArray, factor float64) NDArray {
	return ForEachElement(nd, func(v float64) float64 {
		if v == factor {
			return float64(1)
		} else {
			return float64(0)
		}
	})
}

// ARRAY OPERATIONS

// TODO: Implement this
func ArrayMultiply(a, b NDArray) NDArray {
	if a.Size == 0 || b.Size == 0 {
		return NDArray{}
	}
	panic("numpygo:ArrayMultiply Not implemented")
	return NDArray{}
}

// ARRAY DOT OPERATIONS
func Dot(a, b NDArray, lambda func(float64, float64) float64) NDArray {
	if a.Size == 0 || b.Size == 0 {
		return NDArray{}
	}
	if a.Size != b.Size {
		return NDArray{}
	}
	if !a.Shape.Equals(&b.Shape) {
		return NDArray{}
	}
	newArr := newNDArray(a.DType, a.Shape)
	for i, _ := range newArr.Elements.Values {
		newArr.Elements.Values[i] = lambda(a.Elements.Values[i], b.Elements.Values[i])
	}
	return newArr
}

func Addition(a, b NDArray) NDArray {
	return Dot(a, b, func(x, y float64) float64 {
		return x + y
	})
}

func Subtract(a, b NDArray) NDArray {
	return Dot(a, b, func(x, y float64) float64 {
		return x - y
	})
}

func Multiply(a, b NDArray) NDArray {
	return Dot(a, b, func(x, y float64) float64 {
		return x * y
	})
}

func Division(a, b NDArray) NDArray {
	return Dot(a, b, func(x, y float64) float64 {
		return x / y
	})
}
