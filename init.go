package numpygo

import (
	"math"

	"github.com/praveenpenumaka/numpygo/domain"
	"github.com/praveenpenumaka/numpygo/utils"
)

func newNDArray(dtype string, shape domain.IVector, noinit ...bool) NDArray {
	if dtype == "" {
		return NDArray{}
	}
	if len(shape.Values) == 0 {
		return NDArray{}
	}
	size := shape.Mult()
	elements := domain.Vector{Values: make([]float64, size)}
	dimensions := len(shape.Values)
	strides := utils.GetStrides(size, dimensions, &shape)
	if len(noinit) > 0 {
		if noinit[0] {
			elements = domain.Vector{Values: nil}
		}
	}
	return NDArray{
		Elements: elements,
		Shape:    shape,
		DType:    dtype,
		Size:     size,
		Strides:  strides,
		Dims:     dimensions,
	}
}

func NewNDArray(dtype string, shape []int, noinit ...bool) NDArray {
	if noinit == nil {
		nd := newNDArray(dtype, domain.IVector{Values: shape})
		if nd.Size == 0 {
			return NDArray{}
		}
		nd.Elements.Fill(math.Inf(-1))
		return nd
	} else if noinit[0] {
		nd := newNDArray(dtype, domain.IVector{Values: shape}, noinit[0])
		if nd.Size == 0 {
			return NDArray{}
		}
		return nd
	}
	return NDArray{}
}

func MaxNDArray(dtype string, shape []int) NDArray {
	nd := newNDArray(dtype, domain.IVector{Values: shape})
	if nd.Size == 0 {
		return NDArray{}
	}
	nd.Elements.Fill(math.Inf(1))
	return nd
}

func Zeros(dtype string, shape ...int) NDArray {
	nd := newNDArray(dtype, domain.IVector{Values: shape})
	if nd.Size == 0 {
		return NDArray{}
	}
	nd.Elements.Fill(float64(0))
	return nd
}

func Ones(dtype string, shape ...int) NDArray {
	nd := newNDArray(dtype, domain.IVector{Values: shape})
	if nd.Size == 0 {
		return NDArray{}
	}
	nd.Elements.Fill(float64(1))
	return nd
}

func aRange(start, end, step int) NDArray {
	delta := end - start
	if delta < 0 {
		return NDArray{}
	}
	size := delta / step
	newnd := newNDArray("FLOAT64", domain.IVector{Values: []int{size}})
	for i := range newnd.Elements.Values {
		newnd.Elements.Values[i] = float64(i * step)
	}
	return newnd
}

func Arange(size ...int) NDArray {
	if len(size) <= 0 {
		return NDArray{}
	}
	if len(size) == 1 {
		return aRange(0, size[0], 1)
	}
	if len(size) == 2 {
		return aRange(size[0], size[1], 1)
	}
	return aRange(size[0], size[1], size[2])
}

func (nd *NDArray) CopyLike() NDArray {
	return newNDArray(nd.DType, domain.IVector{Values: nd.Shape.Values})
}
