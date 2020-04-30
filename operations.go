package numpygo

import (
	"errors"
	"github.com/praveenpenumaka/numpygo/domain"
	"github.com/praveenpenumaka/numpygo/utils"
)

func (nd *NDArray) Len() int {
	return nd.Shape.Values[0]
}

func (nd *NDArray) Get(dims *domain.Dimensions) (NDArray, error) {
	//	totalSize := dims.GetTotalSize()
	shape := dims.GetShape()
	if shape == nil {
		return NDArray{}, errors.New("invalid dimensions")
	}
	newnd := newNDArray("FLOAT64", *shape)
	offsetVector := dims.GetStartVector()

	ndIndex := NewNDIndex(shape.Values)
	for vector := ndIndex.Next(); vector != nil; vector = ndIndex.Next() {
		index, err := utils.GetIndexFromVector(vector, &newnd.Strides, &newnd.Shape)
		if err != nil {
			return NDArray{}, err
		}

		vector.Add(offsetVector)
		// Get Index of old array element
		nIndex, iError := utils.GetIndexFromVector(vector, &nd.Strides, &nd.Shape)
		if iError != nil {
			return NDArray{}, iError
		}
		if index >= len(newnd.Elements.Values) || nIndex >= len(nd.Elements.Values) {
			return NDArray{}, errors.New("invalid index for arrays")
		}
		// Copy the element
		newnd.Elements.Values[index] = nd.Elements.Values[nIndex]
	}
	return newnd, nil
}

// TODO: implement this
func (nd *NDArray) GetIndexed(indexes NDArray) NDArray {
	if indexes.Size == 0 {
		return NDArray{}
	}
	if indexes.Dims > 2 {
		return NDArray{}
	}
	if indexes.Dims == 2 {
		if indexes.Shape.Values[1] > 1 && indexes.Shape.Values[0] > 1 {
			return NDArray{}
		}
	}
	if indexes.Size > nd.Shape.Values[0] {
		return NDArray{}
	}

	newShape := &domain.IVector{Values: nil}
	newShape.CopyFrom(nd.Shape.Values)
	newShape.Values[0] = indexes.Shape.Values[0]
	newArray := newNDArray("FLOAT64", *newShape)
	copySize := nd.Strides.Values[0]

	for i, element := range indexes.Elements.Values {
		oldIndex := int(element) * nd.Strides.Values[0]
		newIndex := i * copySize
		for j := 0; j < copySize; j++ {
			newArray.Elements.Values[j+newIndex] = nd.Elements.Values[j+oldIndex]
		}
	}
	return newArray
}

func (nd *NDArray) Set(dims *domain.Dimensions, value float64) error {
	shape := dims.GetShape()
	offsetVector := dims.GetStartVector()
	ndIndex := NewNDIndex(shape.Values)
	for vector := ndIndex.Next(); vector != nil; vector = ndIndex.Next() {
		vector.Add(offsetVector)
		index, err := utils.GetIndexFromVector(vector, &nd.Strides, &nd.Shape)
		if err != nil {
			return err
		}
		// Set the element
		nd.Elements.Values[index] = value
	}
	return nil
}

// TODO: Verify correctness
func (nd *NDArray) Reshape(shape *domain.IVector) bool {
	if shape == nil {
		return false
	}
	size := shape.Mult()
	dimensions := len(shape.Values)
	if size != nd.Size {
		return false
	}
	nd.Size = size
	nd.Shape = *shape
	nd.Strides = utils.GetStrides(size, dimensions, shape)
	nd.Dims = len(nd.Shape.Values)
	return true
}
