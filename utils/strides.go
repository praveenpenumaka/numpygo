package utils

import "github.com/praveenpenumaka/numpygo/domain"

func GetStrides(totalSize, totalDimensions int, shape *domain.IVector) domain.IVector {
	var Strides domain.IVector
	Strides.Values = make([]int, totalDimensions)
	Strides.Zeros()
	for index, dim := range shape.Values {
		rIndex := totalDimensions - index - 1
		totalSize = dim * totalSize
		if index == 0 {
			Strides.Values[rIndex] = 1
		} else {
			Strides.Values[rIndex] = Strides.Values[rIndex+1] * shape.Values[rIndex+1]
		}
	}
	return Strides
}
