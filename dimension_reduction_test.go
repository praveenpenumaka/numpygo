package numpygo

import (
	"github.com/praveenpenumaka/numpygo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArgMaxWithInvalidSize(t *testing.T) {
	arr := NDArray{}
	a := Argmax(arr)
	assert.Equal(t, 0, a.Size)
}

func TestArgMaxWithInvalidAxis(t *testing.T) {
	arr := Rand("FLOAT64", 3, 4, 5)
	a := Argmax(arr, 5)
	assert.Equal(t, 0, a.Size)
}

func TestArgMaxWithoutAxis(t *testing.T) {
	arr := Arange(100)
	arr.Reshape(&domain.IVector{Values: []int{10, 10}})
	a := Argmax(arr)
	assert.Equal(t, 1, a.Size)
	assert.Equal(t, 99.0, a.Elements.Values[0])
}

func TestArgMaxWithFirstAxis(t *testing.T) {
	arr := Arange(6)
	arr.Reshape(&domain.IVector{Values: []int{2, 3}})
	a := Argmax(arr, 0)
	assert.Equal(t, 3, a.Size)
	assert.Equal(t, 1.0, a.Elements.Values[0])
	assert.Equal(t, 1.0, a.Elements.Values[0])
	assert.Equal(t, 1.0, a.Elements.Values[0])
}

func TestArgMinWithInvalidSize(t *testing.T) {
	arr := NDArray{}
	a := Argmin(arr)
	assert.Equal(t, 0, a.Size)
}

func TestArgMinWithInvalidAxis(t *testing.T) {
	arr := Rand("FLOAT64", 3, 4, 5)
	a := Argmin(arr, 5)
	assert.Equal(t, 0, a.Size)
}

func TestArgMinWithoutAxis(t *testing.T) {
	arr := Arange(100)
	arr.Reshape(&domain.IVector{Values: []int{10, 10}})
	a := Argmin(arr)
	assert.Equal(t, 1, a.Size)
	assert.Equal(t, 0.0, a.Elements.Values[0])
}

func TestArgMinWithFirstAxis(t *testing.T) {
	arr := Arange(6)
	arr.Reshape(&domain.IVector{Values: []int{2, 3}})
	a := Argmin(arr, 0)
	assert.Equal(t, 3, a.Size)
	assert.Equal(t, 0.0, a.Elements.Values[0])
	assert.Equal(t, 0.0, a.Elements.Values[0])
	assert.Equal(t, 0.0, a.Elements.Values[0])
}

func TestAminWithoutAxis(t *testing.T) {
	arr := Rand("FLOAT64", 10, 10)
	min := Amin(arr)
	assert.Equal(t, 1, min.Dims)
	assert.Equal(t, 1, min.Shape.Values[0])
}

func TestAminWithoutFirstAxis(t *testing.T) {
	arr := Rand("FLOAT64", 10, 10)
	min := Amin(arr, 0)
	assert.Equal(t, 1, min.Dims)
	assert.Equal(t, 10, min.Shape.Values[0])
}

func TestAmaxNewWithInvalidAxis(t *testing.T) {
	arr := Zeros("FLOAT64", 10, 10)
	dims := []domain.Tuple{domain.Tuple{
		X: 3,
		Y: 3,
	}, domain.Tuple{
		X: 3,
		Y: 3,
	}}
	arr.Set(&domain.Dimensions{Dims: dims}, float64(20))
	maxArr := Amax(arr, 3)
	assert.Equal(t, 0, maxArr.Size)
}

func TestAmaxNewWithAxis(t *testing.T) {
	arr := Zeros("FLOAT64", 10, 10)
	dims := []domain.Tuple{domain.Tuple{
		X: 3,
		Y: 3,
	}, domain.Tuple{
		X: 3,
		Y: 3,
	}}
	arr.Set(&domain.Dimensions{Dims: dims}, float64(20))
	maxArr := Amax(arr, 0)
	assert.Equal(t, float64(20), maxArr.Elements.Values[3])
}

func TestAmaxNewWithoutAxis(t *testing.T) {
	arr := Zeros("FLOAT64", 10, 10)
	dims := []domain.Tuple{domain.Tuple{
		X: 3,
		Y: 3,
	}, domain.Tuple{
		X: 3,
		Y: 3,
	}}
	arr.Set(&domain.Dimensions{Dims: dims}, float64(20))
	wholeArr := Amax(arr)
	assert.Equal(t, float64(20), wholeArr.Elements.Values[0])
}

func TestSumWithEmptyArray(t *testing.T) {
	a := Sum(NDArray{})
	assert.Equal(t, 0, a.Size)
}

func TestSumWithNoAxis(t *testing.T) {
	a := Ones("FLOAT64", 10)
	arr := Sum(a)
	assert.Equal(t, 1, arr.Size)
	assert.Equal(t, 10.0, arr.Elements.Values[0])
}

func TestSumWithInvalidAxis(t *testing.T) {
	a := Zeros("FLOAT64", 10)
	arr := Sum(a, 3)
	assert.Equal(t, 0, arr.Size)
}

func TestSum(t *testing.T) {
	a := Ones("FLOAT64", 10, 10, 10)
	arr := Sum(a, 1)
	assert.Equal(t, 10.0, arr.Elements.Values[0])
	assert.Equal(t, 2, arr.Dims)
	assert.Equal(t, 10, arr.Shape.Values[0])
	assert.Equal(t, 10, arr.Shape.Values[1])
}

///
func TestAmaxWithEmptyArray(t *testing.T) {
	a := Amax(NDArray{})
	assert.Equal(t, 0, a.Size)
}

func TestAmaxWithNoAxis(t *testing.T) {
	a := Ones("FLOAT64", 10)
	arr := Amax(a)
	assert.Equal(t, 1, arr.Size)
	assert.Equal(t, 1.0, arr.Elements.Values[0])
}

func TestAmaxWithInvalidAxis(t *testing.T) {
	a := Zeros("FLOAT64", 10)
	arr := Amax(a, 3)
	assert.Equal(t, 0, arr.Size)
}

func TestAmax(t *testing.T) {
	a := Ones("FLOAT64", 10, 10, 10)
	arr := Amax(a, 1)
	assert.Equal(t, 1.0, arr.Elements.Values[0])
	assert.Equal(t, 2, arr.Dims)
	assert.Equal(t, 10, arr.Shape.Values[0])
	assert.Equal(t, 10, arr.Shape.Values[1])
}

///
func TestAminWithEmptyArray(t *testing.T) {
	a := Amin(NDArray{})
	assert.Equal(t, 0, a.Size)
}

func TestAminWithNoAxis(t *testing.T) {
	a := Ones("FLOAT64", 10)
	arr := Amin(a)
	assert.Equal(t, 1, arr.Size)
	assert.Equal(t, 1.0, arr.Elements.Values[0])
}

func TestAminWithInvalidAxis(t *testing.T) {
	a := Zeros("FLOAT64", 10)
	arr := Amin(a, 3)
	assert.Equal(t, 0, arr.Size)
}

func TestAmin(t *testing.T) {
	a := Ones("FLOAT64", 10, 10, 10)
	arr := Amin(a, 1)
	assert.Equal(t, 1.0, arr.Elements.Values[0])
	assert.Equal(t, 2, arr.Dims)
	assert.Equal(t, 10, arr.Shape.Values[0])
	assert.Equal(t, 10, arr.Shape.Values[1])
}

///
func TestUniqueWithEmptyArray(t *testing.T) {
	a := Unique(NDArray{})
	assert.Equal(t, 0, a.Size)
}

func TestUniqueWithNoAxis(t *testing.T) {
	a := Rand("FLOAT64", 10)
	arr := Unique(a)
	assert.Equal(t, 10, arr.Size)
	assert.NotEqual(t, 1.0, arr.Elements.Values[0])
}

func TestUniqueWithInvalidAxis(t *testing.T) {
	a := Zeros("FLOAT64", 10)
	arr := Unique(a, 3)
	assert.Equal(t, 0, arr.Size)
}
