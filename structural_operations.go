package numpygo

import (
	"fmt"
	"github.com/praveenpenumaka/numpygo/domain"
	"github.com/praveenpenumaka/numpygo/utils"
)

func Split(a *NDArray, index int, axis ...int) (*NDArray, *NDArray) {
	if a == nil {
		return nil, nil
	}
	aAxis := 0
	if axis != nil && len(axis) != 0 {
		if axis[0] >= a.Dims {
			return nil, nil
		}
		aAxis = axis[0]
	}
	if index >= a.Shape.Values[aAxis] {
		return nil, nil
	}

	shapea := domain.IVector{Values: nil}
	shapeb := domain.IVector{Values: nil}
	for i, value := range a.Shape.Values {
		if i == aAxis {
			shapea.Values = append(shapea.Values, index)
			shapeb.Values = append(shapeb.Values, value-index)
		} else {
			shapea.Values = append(shapea.Values, value)
			shapeb.Values = append(shapeb.Values, value)
		}
	}
	sa := newNDArray("FLOAT64", shapea)
	sb := newNDArray("FLOAT64", shapeb)

	ndIndex := NewNDIndex(a.Shape.Values)
	for vector := ndIndex.Next(); vector != nil; vector = ndIndex.Next() {
		indexAlongAxis := vector.Values[aAxis]
		diff := indexAlongAxis - index
		index, err := utils.GetIndexFromVector(vector, &a.Strides, &a.Shape)
		if err != nil {
			return nil, nil
		}
		value := a.Elements.Values[index]
		if diff >= 0 {
			vector.Values[aAxis] = diff
			newIndex, err := utils.GetIndexFromVector(vector, &sb.Strides, &sb.Shape)
			if err != nil {
				return nil, nil
			}
			sb.Elements.Values[newIndex] = value
		} else {
			newIndex, err := utils.GetIndexFromVector(vector, &sa.Strides, &sa.Shape)
			if err != nil {
				return nil, nil
			}
			sa.Elements.Values[newIndex] = value
		}
	}
	return sa, sb
}

func Concatenate(a *NDArray, b *NDArray, axis ...int) *NDArray {
	if a == nil || b == nil {
		return nil
	}
	aAxis := 0
	if axis != nil && len(axis) != 0 {
		if axis[0] > a.Dims {
			return nil
		}
		aAxis = axis[0]
	}
	if a.Dims != b.Dims {
		// Cannot append if arrays has different dimensions
		return nil
	}

	// All the axis should be same except for along axis
	aShape := a.Shape.Remove(aAxis)
	bShape := b.Shape.Remove(aAxis)
	if !aShape.Equals(bShape) {
		return nil
	}

	newShape := domain.IVector{Values: nil}
	for i, value := range a.Shape.Values {
		if i == aAxis {
			newShape.Values = append(newShape.Values, b.Shape.Values[aAxis]+a.Shape.Values[aAxis])
		} else {
			newShape.Values = append(newShape.Values, value)
		}
	}

	cArray := newNDArray("FLOAT64", newShape)

	ndIndex := NewNDIndex(newShape.Values)
	for vector := ndIndex.Next(); vector != nil; vector = ndIndex.Next() {
		newIndex, err := utils.GetIndexFromVector(vector, &cArray.Strides, &cArray.Shape)
		if err != nil {
			return nil
		}
		value := float64(0)
		diff := vector.Values[aAxis] - a.Shape.Values[aAxis]
		if diff >= 0 {
			vector.Values[aAxis] = diff
			index, err := utils.GetIndexFromVector(vector, &b.Strides, &b.Shape)
			if err != nil {
				return nil
			}
			value = b.Elements.Values[index]
		} else {
			index, err := utils.GetIndexFromVector(vector, &a.Strides, &a.Shape)
			if err != nil {
				return nil
			}
			value = a.Elements.Values[index]
		}
		cArray.Elements.Values[newIndex] = value
	}

	return cArray
}

func FilterRows(a *NDArray,lambda func(row *domain.Vector) bool) *NDArray {
	newArray := &NDArray{
		Elements: domain.Vector{},
		Shape:    domain.IVector{},
		Strides:  domain.IVector{},
		DType:    "FLOAT64",
		Dims:     a.Dims,
		Size:     0,
	}
	newArray.Shape.CopyFrom(a.Shape.Values)
	newArray.Shape.Zeros()
	newArray.Strides.CopyFrom(a.Strides.Values)
	size :=0
	count :=0
	for i :=0;i<a.Shape.Values[0];i++ {
		dim := ParseDim(fmt.Sprintf("%d,:",i),&a.Shape)
		row,rowErr := a.Get(dim)
		if rowErr == nil{
			if lambda(&row.Elements){
				newArray.Elements.Values = append(newArray.Elements.Values, row.Elements.Values...)
				count++
				size = size + len(row.Elements.Values)
				newArray.Shape.Values[0] = count
				newArray.Shape.Values[1] = row.Size
				newArray.Size= size
			}
		}
	}
	return newArray
}

func Diag(nd *NDArray, k ...int) *NDArray{
	kLen := 0
	if len(k) > 0 {
		kLen = k[0]
	}
	if nd.Dims != 2 {
		return nil
	}
	axis := nd.Shape.Min()
	newArr:=NewNDArray("FLOAT64",[]int{axis,1},true)
	size:=0
	for i:=0;i<axis;i++{
		x:=i
		y:=i+kLen
		oldIndex := nd.Strides.Values[0]*x+nd.Strides.Values[1]*y
		if x>=0 && y>=0 && oldIndex >=0 && oldIndex  < nd.Size{
			newArr.Elements.Values = append(newArr.Elements.Values, nd.Elements.Values[oldIndex])
			size++
		}
	}
	newArr.Shape.Values[0] = size
	newArr.Strides.Values[0] = 1
	newArr.Size = size
	return newArr
}

// TODO: Implement this
func Transpose(nd *NDArray) *NDArray{
	return nd
}
