package numpygo

import (
	"github.com/praveenpenumaka/numpygo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDimWithNilShape(t *testing.T) {
	dims := ParseDim(":", nil)
	assert.Nil(t, dims)
}

func TestParseDimWithEmptyShape(t *testing.T) {
	shape := &domain.IVector{Values: []int{0, 0}}
	dims := ParseDim(":", shape)
	assert.Nil(t, dims)
}

func TestParseDimWithOneDShape(t *testing.T) {
	shapea := &domain.IVector{Values: []int{10, 1}}
	shapeb := &domain.IVector{Values: []int{1, 10}}
	dimsa := ParseDim(":", shapea)
	dimsb := ParseDim(":", shapeb)
	assert.NotNil(t, dimsa)
	assert.NotNil(t, dimsb)
}

func TestParseDimWithOneDShapeHasTotalSize(t *testing.T) {
	shapea := &domain.IVector{Values: []int{10, 1}}
	shapeb := &domain.IVector{Values: []int{1, 10}}
	dimsa := ParseDim(":", shapea)
	dimsb := ParseDim(":", shapeb)
	assert.Equal(t, 2, len(dimsa.Dims))
	assert.Equal(t, 0, dimsa.Dims[0].X)
	assert.Equal(t, 9, dimsa.Dims[0].Y)
	assert.Equal(t, 0, dimsa.Dims[1].X)
	assert.Equal(t, 0, dimsa.Dims[1].Y)
	assert.Equal(t, 2, len(dimsb.Dims))
	assert.Equal(t, 0, dimsb.Dims[0].X)
	assert.Equal(t, 0, dimsb.Dims[0].Y)
	assert.Equal(t, 0, dimsb.Dims[1].X)
	assert.Equal(t, 9, dimsb.Dims[1].Y)
}

func TestParseDimWithSingleColonFirst(t *testing.T) {
	shape := &domain.IVector{Values: []int{10, 8}}
	dims := ParseDim(":,3", shape)
	assert.NotNil(t, dims)
	assert.Equal(t, 2, len(dims.Dims))
	assert.Equal(t, 0, dims.Dims[0].X)
	assert.Equal(t, 9, dims.Dims[0].Y)
	assert.Equal(t, 3, dims.Dims[1].X)
	assert.Equal(t, 3, dims.Dims[1].Y)
}

func TestParseDimWithSingleColonSecond(t *testing.T) {
	shape := &domain.IVector{Values: []int{10, 8}}
	dims := ParseDim("3,:", shape)
	assert.NotNil(t, dims)
	assert.Equal(t, 2, len(dims.Dims))
	assert.Equal(t, 3, dims.Dims[0].X)
	assert.Equal(t, 3, dims.Dims[0].Y)
	assert.Equal(t, 0, dims.Dims[1].X)
	assert.Equal(t, 7, dims.Dims[1].Y)
}

//
func TestParseDimWithTwoColonFirst(t *testing.T) {
	shape := &domain.IVector{Values: []int{10, 8}}
	dims := ParseDim(":,:3", shape)
	assert.NotNil(t, dims)
	assert.Equal(t, 2, len(dims.Dims))
	assert.Equal(t, 0, dims.Dims[0].X)
	assert.Equal(t, 9, dims.Dims[0].Y)
	assert.Equal(t, 0, dims.Dims[1].X)
	assert.Equal(t, 3, dims.Dims[1].Y)

	dims = ParseDim(":,3:", shape)
	assert.NotNil(t, dims)
	assert.Equal(t, 2, len(dims.Dims))
	assert.Equal(t, 0, dims.Dims[0].X)
	assert.Equal(t, 9, dims.Dims[0].Y)
	assert.Equal(t, 3, dims.Dims[1].X)
	assert.Equal(t, 7, dims.Dims[1].Y)

}

func TestParseDimWithTwoColonSecond(t *testing.T) {
	shape := &domain.IVector{Values: []int{10, 8}}
	dims := ParseDim(":3,:", shape)
	assert.NotNil(t, dims)
	assert.Equal(t, 2, len(dims.Dims))
	assert.Equal(t, 0, dims.Dims[0].X)
	assert.Equal(t, 3, dims.Dims[0].Y)
	assert.Equal(t, 0, dims.Dims[1].X)
	assert.Equal(t, 7, dims.Dims[1].Y)

	dims = ParseDim("3:,:", shape)
	assert.NotNil(t, dims)
	assert.Equal(t, 2, len(dims.Dims))
	assert.Equal(t, 3, dims.Dims[0].X)
	assert.Equal(t, 9, dims.Dims[0].Y)
	assert.Equal(t, 0, dims.Dims[1].X)
	assert.Equal(t, 7, dims.Dims[1].Y)
}

func TestGetDimWithEmptyString(t *testing.T) {
	shape := &domain.IVector{Values: []int{10, 10}}
	dim := ParseDim("", shape)
	assert.Nil(t, dim)
}

func TestGetDimForAllAndInvalidShape(t *testing.T) {
	dim := ParseDim(":", nil)
	assert.Nil(t, dim)
}

func TestGetDimForAllAndValidShape(t *testing.T) {
	dim := ParseDim(":", &domain.IVector{Values: []int{10, 10}})
	assert.NotNil(t, dim)
}
