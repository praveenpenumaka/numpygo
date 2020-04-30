package utils

import (
	"errors"
	"github.com/praveenpenumaka/numpygo/domain"
)

func GetAxis(axis []int) int {
	if axis != nil && len(axis) != 0 {
		return axis[0]
	}
	return 0
}

func GetIVector(newIndex int, newStrides, newShape []int) (*domain.IVector, error) {
	var vector []int
	reminder := newIndex
	checkIndex := 0
	for sIndex, stride := range newStrides {
		mod := int(reminder / stride)
		if mod >= newShape[sIndex] {
			return nil, errors.New("invalid index")
		}
		reminder = reminder % stride
		checkIndex += stride * mod
		vector = append(vector, mod)
	}
	if checkIndex != newIndex {
		return nil, errors.New("invalid metadata")
	}
	return &domain.IVector{Values: vector}, nil
}

func GetVector(newIndex int, newStrides, newShape []int) ([]int, error) {
	var vector []int
	reminder := newIndex
	checkIndex := 0
	for sIndex, stride := range newStrides {
		mod := int(reminder / stride)
		if mod >= newShape[sIndex] {
			return nil, errors.New("invalid index")
		}
		reminder = reminder % stride
		checkIndex += stride * mod
		vector = append(vector, mod)
	}
	if checkIndex != newIndex {
		return nil, errors.New("invalid metadata")
	}
	return vector, nil
}

func GetVectors(indexes []int, newStrides, newShape []int) ([][]int, error) {
	var vectors [][]int
	for _, index := range indexes {
		vec, err := GetVector(index, newStrides, newShape)
		if err != nil {
			return vectors, err
		}
		vectors = append(vectors, vec)
	}
	return vectors, nil
}

func GetIndex(vector, newStrides, size []int) (int, error) {
	index := 0
	for sIndex, stride := range newStrides {
		if vector[sIndex] >= size[sIndex] {
			return 0, errors.New("invalid vector")
		}
		index += stride * vector[sIndex]
	}
	return index, nil
}

func GetIndexFromVector(vector *domain.IVector, newStrides, shape *domain.IVector) (int, error) {
	if vector == nil || vector.Values == nil {
		return 0, errors.New("empty vector")
	}
	index := 0
	for sIndex, stride := range newStrides.Values {
		if vector.Values[sIndex] >= shape.Values[sIndex] {
			return 0, errors.New("invalid vector")
		}
		index += stride * vector.Values[sIndex]
	}
	return index, nil
}

func GetIndexes(vectors [][]int, newStrides, newShape []int) ([]int, error) {
	var indexes []int
	for _, vector := range vectors {
		in, err := GetIndex(vector, newStrides, newShape)
		if err != nil {
			return []int{}, err
		}
		indexes = append(indexes, in)
	}
	return indexes, nil
}

func GetShapeFromDims(dims []domain.Tuple) []int {
	var shape []int
	for _, dim := range dims {
		if dim.IsValidDimension() {
			size := 1
			if dim.X != dim.Y {
				size = dim.Y - dim.X
			}
			shape = append(shape, size)
		}
	}
	return shape
}

// Converts dimension tuples to single vector
func DimsToVector(dims []domain.Tuple) []int {
	var vector []int
	for _, dim := range dims {
		vector = append(vector, dim.X)
	}
	return vector
}

func RemoveItemInSlice(a []int, index int) []int {
	var newA []int
	for i, i2 := range a {
		if i != index {
			newA = append(newA, i2)
		}
	}
	return newA
}
