package numpygo

import (
	"github.com/praveenpenumaka/numpygo/domain"
	"github.com/praveenpenumaka/numpygo/utils"
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
	nd.Elements.Exp()
	return nd
}

func Log2(nd NDArray) NDArray {
	nd.Elements.Log2()
	return nd
}

func Log(nd NDArray) NDArray {
	nd.Elements.Log()
	return nd
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

// TODO: Support broadcasting
//In order to broadcast, the size of the trailing axes for both arrays in an
// operation must either be the same size or one of them must be one.

//The size of the result array created by broadcast operations is the maximum size
// along each dimension from the input arrays.
func IsDotOperationAllowed(a, b NDArray, boardcasting ...bool) bool {
	broadcastingEnabled := len(boardcasting) > 0
	if (a.Size == 0 || b.Size == 0) && !broadcastingEnabled {
		return false
	}
	if a.Size != b.Size && !broadcastingEnabled {
		return false
	}
	if !a.Shape.Equals(&b.Shape) && !broadcastingEnabled {
		return false
	}
	if broadcastingEnabled {

	}
	return true
}

func getDotShape(a, b domain.IVector) (domain.IVector, string) {
	adims := len(a.Values)
	bdims := len(b.Values)
	if adims == 0 || bdims == 0 {
		return domain.IVector{}, ""
	}
	da := &a
	db := &b
	if da.Equals(db) {
		return a, "a"
	}
	dims := adims
	largeVector := ""
	if adims < bdims {
		dims = bdims
		pada := bdims - adims
		da.Padding("LEAD", 1, pada)
	} else if adims > bdims {
		padb := adims - bdims
		db.Padding("LEAD", 1, padb)
	}
	newShape := domain.IVector{Values: make([]int, dims)}
	for i := 0; i < dims; i++ {
		d1 := a.Values[i]
		d2 := b.Values[i]
		if d1 == d2 {
			newShape.Values[i] = d1
		} else if d1 == 1 || d2 == 1 {
			if d1 > d2 {
				if largeVector == "" {
					largeVector = "a"
				} else if largeVector != "a" {
					return domain.IVector{}, ""
				}
				newShape.Values[i] = d1
			} else if d2 > d1 {
				if largeVector == "" {
					largeVector = "b"
				} else if largeVector != "b" {
					return domain.IVector{}, ""
				}
				newShape.Values[i] = d2
			}
		} else {
			return domain.IVector{}, ""
		}
	}
	return newShape, largeVector
}

func broadcastTo(a NDArray, vector domain.IVector, broadcasting bool) (NDArray, error) {
	if !broadcasting {
		return a, nil
	}
	newArray := newNDArray(a.DType, vector)
	if a.Dims == 1 {
		for i := 0; i < newArray.Size; i = i + a.Size {
			newArray.Elements.FillIndex(a.Elements.Values, i)
		}
		return newArray, nil
	} else if a.Dims >= 2 {
		iter := NewNDIndex(newArray.Shape.Values)
		bStrides := &a.Strides
		bStrides.Padding("LEAD", a.Strides.Values[a.Dims-1], len(vector.Values)-a.Dims)
		//bStrides := domain.IVector{Values:[]int{1,1,1}}
		bShape := &a.Shape
		bShape.Padding("LEAD", 1, len(vector.Values)-a.Dims)
		//bShape := domain.IVector{Values:[]int{1,3,3}}
		i := 0
		for vector := iter.Next(); vector != nil; vector = iter.Next() {
			oldIndex, err := utils.GetIndexFromVector(vector, bStrides, bShape, true)
			if err != nil {
				panic("Invalid index for broadcasting")
			}
			newArray.Elements.Values[i] = a.Elements.Values[oldIndex]
			i++
		}
		return newArray, nil
	}
	return newArray, nil
}

// ARRAY DOT OPERATIONS
func Dot(x, y NDArray, lambda func(float64, float64) float64) NDArray {
	if x.Size == 0 || y.Size == 0 {
		return NDArray{}
	}
	a := x
	b := y
	newShape, largeVector := getDotShape(x.Shape, y.Shape)
	if len(newShape.Values) == 0 || largeVector == "" {
		return NDArray{}
	}
	if largeVector == "b" {
		a = y
		b = x
	}
	newArr := newNDArray(a.DType, newShape)
	iter := NewNDIndex(newArr.Shape.Values)
	bStrides := &b.Strides
	bStrides.Padding("LEAD", b.Strides.Values[b.Dims-1], a.Dims-b.Dims)
	bShape := &b.Shape
	bShape.Padding("LEAD", 1, a.Dims-b.Dims)

	i := 0
	for vector := iter.Next(); vector != nil; vector = iter.Next() {
		bIndex, err := utils.GetIndexFromVector(vector, bStrides, bShape, true)
		if err != nil {
			return NDArray{}
		}
		newArr.Elements.Values[i] = lambda(a.Elements.Values[i], b.Elements.Values[bIndex])
		i++
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
