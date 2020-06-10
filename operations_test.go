package numpygo

import (
	"testing"

	"github.com/praveenpenumaka/numpygo/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetWithFloat64(t *testing.T) {
	array := Ones("FLOAT64", 10, 10)
	dims := []domain.Tuple{
		{
			X: 0,
			Y: 9,
		},
		{
			X: 0,
			Y: 9,
		},
	}
	newArray, _ := array.Get(&domain.Dimensions{Dims: dims})
	assert.NotNil(t, newArray)
	assert.Equal(t, float64(1), newArray.Elements.Values[0])
	assert.Equal(t, 100, len(newArray.Elements.Values))
	assert.Equal(t, 100, newArray.Size)
	assert.Equal(t, 2, len(newArray.Shape.Values))
	assert.Equal(t, 10, newArray.Shape.Values[0])
	assert.Equal(t, 10, newArray.Shape.Values[1])
}

func TestGetPartialData(t *testing.T) {
	array := Ones("FLOAT64", 10, 10)
	dims := []domain.Tuple{
		{X: 1, Y: 2},
		{X: 1, Y: 2},
	}
	newArray, _ := array.Get(&domain.Dimensions{Dims: dims})
	assert.Equal(t, float64(1), newArray.Elements.Values[0])
	assert.Equal(t, 4, len(newArray.Elements.Values))
	assert.Equal(t, 4, newArray.Size)
	assert.Equal(t, 2, len(newArray.Shape.Values))
	assert.Equal(t, 2, newArray.Shape.Values[0])
	assert.Equal(t, 2, newArray.Shape.Values[1])
}

func TestGetIncorrectPartialData(t *testing.T) {
	array := Ones("FLOAT64", 10, 10)
	dims := []domain.Tuple{
		{X: 0, Y: 0},
		{X: 1, Y: 10},
	}
	newArray, err := array.Get(&domain.Dimensions{Dims: dims})
	assert.NotNil(t, err)
	assert.NotNil(t, newArray)
	assert.Equal(t, 0, newArray.Size)
	dims = []domain.Tuple{
		{X: -1, Y: 1},
		{X: 1, Y: 1},
	}
	newArray, err = array.Get(&domain.Dimensions{Dims: dims})
	assert.NotNil(t, err)
	assert.NotNil(t, newArray)
	assert.Equal(t, 0, newArray.Size)
}

func TestGetCorrectnessInOneD(t *testing.T) {
	array := Ones("FLOAT64", 10)
	dims := []domain.Tuple{
		{X: 0, Y: 3},
	}
	newArray, err := array.Get(&domain.Dimensions{Dims: dims})
	assert.Nil(t, err)
	assert.NotNil(t, newArray)
	assert.Equal(t, 4, newArray.Shape.Values[0])
}

func TestGetIndexedWithNilIndexes(t *testing.T) {
	ones := Ones("FLOAT64", 10, 8)
	getIn := ones.GetIndexed(NDArray{})
	assert.NotNil(t, getIn)
	assert.Equal(t, 0, getIn.Size)
}

func TestGetIndexedWithInvalidIndexes(t *testing.T) {
	ones := Ones("FLOAT64", 10, 8)
	aind := Arange(11)
	getIn := ones.GetIndexed(aind)
	assert.NotNil(t, getIn)
	assert.Equal(t, 0, getIn.Size)
}

func TestGetIndexedWith2DIndexes(t *testing.T) {
	ones := Ones("FLOAT64", 10, 8)
	aind := Ones("FLOAT64", 10, 10)
	getIn := ones.GetIndexed(aind)
	assert.NotNil(t, getIn)
	assert.Equal(t, 0, getIn.Size)
}

func TestGetIndexedWithValidIndexes(t *testing.T) {
	ones := Ones("FLOAT64", 10, 8)
	aind := Arange(10)
	getIn := ones.GetIndexed(aind)
	assert.NotNil(t, getIn)
	assert.Equal(t, 10, getIn.Len())
	assert.Equal(t, 80, getIn.Size)
	for i := 0; i < getIn.Size; i++ {
		assert.Equal(t, float64(1), getIn.Elements.Values[i])
	}
}

func TestArraySet(t *testing.T) {
	arr := Zeros("FLOAT64", 10)
	dims := []domain.Tuple{
		{X: 0, Y: 0},
	}
	assert.NoError(t, arr.Set(&domain.Dimensions{Dims: dims}, float64(1)))
	assert.Equal(t, float64(1), arr.Elements.Values[0])
	dims = []domain.Tuple{
		{X: 1, Y: 9},
	}
	assert.NoError(t, arr.Set(&domain.Dimensions{Dims: dims}, float64(1)))
	assert.Equal(t, float64(1), arr.Elements.Values[1])
}

func TestReshapeWithInvalidShape(t *testing.T) {
	array := Ones("FLOAT64", 10, 10)
	shape := &domain.IVector{Values: []int{1, 2}}
	isSuccess := array.Reshape(shape)
	assert.False(t, isSuccess)
}

func TestReshapeWithInvalidDimensions(t *testing.T) {
	array := Ones("FLOAT64", 10, 10)
	shape := &domain.IVector{Values: []int{1, 2, 3}}
	isSuccess := array.Reshape(shape)
	assert.False(t, isSuccess)
}

func TestReshapeWithNilShape(t *testing.T) {
	array := Ones("FLOAT64", 10, 10)
	isSuccess := array.Reshape(nil)
	assert.False(t, isSuccess)
}

func TestReshapeWithCorrectShape(t *testing.T) {
	array := Ones("FLOAT64", 10, 10)
	shape := &domain.IVector{Values: []int{1, 100}}
	isSuccess := array.Reshape(shape)
	assert.True(t, isSuccess)
	assert.Equal(t, 1, array.Shape.Values[0])
	assert.Equal(t, 100, array.Shape.Values[1])
	assert.Equal(t, 100, array.Size)
}
