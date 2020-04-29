package numpygo

import (
	"github.com/praveenpenumaka/numpygo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitNilArray(t *testing.T) {
	sa, sb := Split(nil, 2)
	assert.Nil(t, sa)
	assert.Nil(t, sb)
}

func TestSplitWithInvalidIndex(t *testing.T) {
	a := Ones("FLOAT64", 10, 10)
	sa, sb := Split(a, 12)
	assert.Nil(t, sa)
	assert.Nil(t, sb)
}

func TestSplitWithInvalidAxis(t *testing.T) {
	a := Ones("FLOAT64", 10, 10)
	sa, sb := Split(a, 5, 2)
	assert.Nil(t, sa)
	assert.Nil(t, sb)
}

func TestSplitArray(t *testing.T) {
	a := Ones("FLOAT64", 1, 2)
	b := Zeros("FLOAT64", 2, 2)
	c := Concatenate(a, b)
	sa, sb := Split(c, 1, 0)
	assert.NotNil(t, sa)
	assert.NotNil(t, sb)
	assert.Equal(t, 1, sa.Shape.Values[0])
	assert.Equal(t, 2, sa.Shape.Values[1])
	for _, value := range sa.Elements.Values {
		assert.Equal(t, float64(1), value)
	}
	assert.Equal(t, 2, sb.Shape.Values[0])
	assert.Equal(t, 2, sb.Shape.Values[1])
	for _, value := range sb.Elements.Values {
		assert.Equal(t, float64(0), value)
	}
}

func TestConcatenateWithNilArrays(t *testing.T) {
	c := Concatenate(nil, nil)
	assert.Nil(t, c)
}

func TestConcatenateWithInvalidAxis(t *testing.T) {
	a := Ones("FLOAT64", 10, 10)
	b := Ones("FLOAT64", 10, 10, 10)
	c := Concatenate(a, b, 4)
	assert.Nil(t, c)
}

func TestConcatenateWithDifferentDimensionArrays(t *testing.T) {
	a := Ones("FLOAT64", 10, 10)
	b := Ones("FLOAT64", 10, 10, 10)
	c := Concatenate(a, b)
	assert.Nil(t, c)
}

func TestConcatenateWithNonMatchingDimensions(t *testing.T) {
	a := Ones("FLOAT64", 3, 9)
	b := Ones("FLOAT64", 3, 10)
	c := Concatenate(a, b)
	assert.Nil(t, c)
}

func TestConcatenate(t *testing.T) {
	a := Ones("FLOAT64", 1, 2)
	b := Zeros("FLOAT64", 2, 2)
	c := Concatenate(a, b)
	assert.NotNil(t, c)
	for i, value := range c.Elements.Values {
		if i <= 1 {
			assert.Equal(t, float64(1), value)
		} else {
			assert.Equal(t, float64(0), value)
		}
	}
}


func TestDiagWithInvalidDims(t *testing.T)  {
	a := Ones("FLOAT64",2,2,2)
	b:= Diag(a)
	assert.Nil(t,b)
}


func TestDiag(t *testing.T)  {
	a := Arange(9)
	a.Reshape(&domain.IVector{Values:[]int{3,3}})
	b:= Diag(a)
	assert.NotNil(t,b)
	assert.Equal(t,3,b.Shape.Values[0])
	assert.Equal(t,1,b.Shape.Values[1])
	assert.Equal(t,3,b.Size)
	assert.Equal(t,float64(0),b.Elements.Values[0])
	assert.Equal(t,float64(4),b.Elements.Values[1])
	assert.Equal(t,float64(8),b.Elements.Values[2])
}

func TestDiagWithPositiveK(t *testing.T)  {
	a := Arange(9)
	a.Reshape(&domain.IVector{Values:[]int{3,3}})
	b:= Diag(a,1)
	assert.NotNil(t,b)
	assert.Equal(t,2,b.Shape.Values[0])
	assert.Equal(t,1,b.Shape.Values[1])
	assert.Equal(t,2,b.Size)
	assert.Equal(t,float64(1),b.Elements.Values[0])
	assert.Equal(t,float64(5),b.Elements.Values[1])
}

func TestDiagWithNegativeK(t *testing.T)  {
	a := Arange(9)
	a.Reshape(&domain.IVector{Values:[]int{3,3}})
	b:= Diag(a,-2)
	assert.NotNil(t,b)
	assert.Equal(t,1,b.Shape.Values[0])
	assert.Equal(t,1,b.Shape.Values[1])
	assert.Equal(t,1,b.Size)
	assert.Equal(t,float64(6),b.Elements.Values[0])
}
