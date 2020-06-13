package numpygo

import (
	"github.com/praveenpenumaka/numpygo/domain"
	"github.com/praveenpenumaka/numpygo/utils"
)

type NDArray struct {
	Elements domain.Vector
	Shape    domain.IVector
	Strides  domain.IVector
	DType    string
	Dims     int
	Size     int
}

// TODO: Write test cases
func (nd *NDArray) Mean(axis ...int) NDArray {
	if len(axis) == 0 {
		newArray := newNDArray(nd.DType, domain.IVector{Values: []int{1}})
		newArray.Elements.Values[0] = nd.Elements.Mean()
		return newArray
	}

	newShape := nd.Shape.Remove(axis[0])
	ndIndex := NewNDIndex(nd.Shape.Values)
	newArray := MaxNDArray(nd.DType, newShape.Values)
	counter := Zeros("FLOAT64", newArray.Size)
	for vector := ndIndex.Next(); vector != nil; vector = ndIndex.Next() {
		oldIndex, err := utils.GetIndexFromVector(vector, &nd.Strides, &nd.Shape)
		if err != nil {
			return NDArray{}
		}
		newVector := vector.Remove(axis[0])
		newIndex, err := utils.GetIndexFromVector(newVector, &newArray.Strides, &newArray.Shape)
		if err != nil {
			return NDArray{}
		}
		newArray.Elements.Values[newIndex] = newArray.Elements.Values[newIndex] + nd.Elements.Values[oldIndex]
		counter.Elements.Values[newIndex]++
	}
	for i, value := range newArray.Elements.Values {
		count := counter.Elements.Values[i]
		newArray.Elements.Values[i] = value / count
	}
	return newArray
}

// TODO: Add test cases
func Equals(a NDArray, b NDArray) NDArray {
	if a.Size == 0 || b.Size == 0 {
		return NDArray{}
	}
	if !a.Shape.Equals(&b.Shape) {
		return NDArray{}
	}
	newShape := domain.IVector{Values: nil}
	newShape.CopyFrom(a.Shape.Values)
	newArray := newNDArray("FLOAT64", newShape)
	for i, value := range a.Elements.Values {
		if value == b.Elements.Values[i] {
			newArray.Elements.Values[i] = 1
		} else {
			newArray.Elements.Values[i] = 0
		}
	}
	return newArray
}

// Clip (limit) the values in an array.
func Clip(a NDArray, min, max float64) NDArray {
	if a.Size == 0 {
		return NDArray{}
	}
	vec := &a.Elements
	vec.Clip(min, max)
	a.Elements = *vec
	return a
}
