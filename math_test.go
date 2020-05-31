package numpygo

import (
	"github.com/praveenpenumaka/numpygo/domain"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMathLambdaExp(t *testing.T) {
	ones := Ones("FLOAT64", 10, 10, 10)
	logOnes := Exp(ones)
	assert.Equal(t, math.Exp(1), logOnes.Elements.Values[0])
}

func TestMathLambdaLog2(t *testing.T) {
	ones := Ones("FLOAT64", 10, 10, 10)
	logOnes := Log2(ones)
	assert.Equal(t, math.Log2(1), logOnes.Elements.Values[0])
}

func TestMathLambdaLog(t *testing.T) {
	ones := Ones("FLOAT64", 10, 10, 10)
	logOnes := Log(ones)
	assert.Equal(t, float64(0), logOnes.Elements.Values[0])
}

func TestMathLambdaMult(t *testing.T) {
	ones := Ones("FLOAT64", 10, 10, 10)
	logOnes := Mult(ones, 5.2)
	assert.Equal(t, 5.2, logOnes.Elements.Values[0])
}

func TestMathLambdaDiv(t *testing.T) {
	ones := Ones("FLOAT64", 10, 10, 10)
	logOnes := Div(ones, 0.5)
	assert.Equal(t, float64(2), logOnes.Elements.Values[0])
}

func TestMathLambdaDivByZero(t *testing.T) {
	ones := Ones("FLOAT64", 10, 10, 10)
	logOnes := Div(ones, 0)
	assert.Equal(t, math.Inf(1), logOnes.Elements.Values[0])
}

func TestMathLambdaPow(t *testing.T) {
	ones := Ones("FLOAT64", 10, 10, 10)
	logOnes := Pow(ones, 10)
	assert.Equal(t, float64(1), logOnes.Elements.Values[0])
}

func TestMathLambdaEqual(t *testing.T) {
	ones := Ones("FLOAT64", 10, 10, 10)
	logOnes := Equal(ones, 0.5)
	assert.Equal(t, float64(0), logOnes.Elements.Values[0])
	logOnes = Equal(ones, float64(1))
	assert.Equal(t, float64(1), logOnes.Elements.Values[0])
}

func TestAdditionFailForNilInput(t *testing.T) {
	a := Addition(NDArray{}, NDArray{})
	assert.NotNil(t, a)
	assert.Equal(t, 0, a.Size)
}

/*
func TestAdditionFailForIncompatibleSize(t *testing.T) {
	a := Ones("FLOAT64", 10, 5)
	b := Ones("FLOAT64", 10, 6)
	c := Addition(a, b)
	assert.NotNil(t, c)
	assert.Equal(t, 0, c.Size)
}
*/

func TestAdditionSuccess(t *testing.T) {
	a := Ones("FLOAT64", 10, 5)
	b := Ones("FLOAT64", 10, 5)
	c := Addition(a, b)
	assert.Equal(t, int(10), c.Shape.Values[0])
	assert.Equal(t, int(5), c.Shape.Values[1])
	assert.Equal(t, float64(2), c.Elements.Values[0])
}

func TestSubtractFailForNilInput(t *testing.T) {
	a := Subtract(NDArray{}, NDArray{})
	assert.NotNil(t, a)
	assert.Equal(t, 0, a.Size)
}

func TestSubtractSuccess(t *testing.T) {
	a := Ones("FLOAT64", 10, 5)
	b := Ones("FLOAT64", 10, 5)
	c := Subtract(a, b)
	assert.Equal(t, int(10), c.Shape.Values[0])
	assert.Equal(t, int(5), c.Shape.Values[1])
	assert.Equal(t, float64(0), c.Elements.Values[0])
}

func TestMultiplyFailForNilInput(t *testing.T) {
	a := Multiply(NDArray{}, NDArray{})
	assert.NotNil(t, a)
	assert.Equal(t, 0, a.Size)
}

func TestMultiplySuccess(t *testing.T) {
	a := Ones("FLOAT64", 10, 5)
	b := Ones("FLOAT64", 10, 5)
	c := Multiply(a, b)
	assert.Equal(t, int(10), c.Shape.Values[0])
	assert.Equal(t, int(5), c.Shape.Values[1])
	assert.Equal(t, float64(1), c.Elements.Values[0])
}

func TestDivisionFailForNilInput(t *testing.T) {
	a := Division(NDArray{}, NDArray{})
	assert.NotNil(t, a)
	assert.Equal(t, 0, a.Size)
}

func TestDivisionSuccess(t *testing.T) {
	a := Ones("FLOAT64", 10, 5)
	b := Ones("FLOAT64", 10, 5)
	c := Division(a, b)
	assert.Equal(t, int(10), c.Shape.Values[0])
	assert.Equal(t, int(5), c.Shape.Values[1])
	assert.Equal(t, float64(1), c.Elements.Values[0])
}

// Broadcasting Related

func TestBroadcastToWithoutBroadcasting(t *testing.T) {
	a := Ones("FLOAT64", 3, 3, 3)
	b, err := broadcastTo(a, domain.IVector{Values: []int{3, 3, 3}}, false)
	assert.Nil(t, err)
	assert.Equal(t, a.Dims, b.Dims)
	assert.True(t, a.Shape.Equals(&b.Shape))
}

func TestBroadcastToWithBroadcastingSingleDim(t *testing.T) {
	a := Ones("FLOAT64", 1)
	a.Elements.Values[0] = 5.0
	b, err := broadcastTo(a, domain.IVector{Values: []int{3, 3, 3}}, true)
	assert.Nil(t, err)
	assert.Equal(t, 3, b.Dims)
	for _, v := range b.Elements.Values {
		assert.Equal(t, float64(5), v)
	}
}

func TestBroadcastToWithBroadcastingTwoDimsWithoutOnes(t *testing.T) {
	a := Arange(9)
	a.Reshape(&domain.IVector{Values: []int{3, 3}})
	b, err := broadcastTo(a, domain.IVector{Values: []int{3, 3, 3}}, true)
	assert.Nil(t, err)
	assert.Equal(t, 3, b.Dims)
	expected := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 0, 1, 2, 3, 4, 5, 6, 7, 8, 0, 1, 2, 3,
		4, 5, 6, 7, 8}
	for i, v := range b.Elements.Values {
		assert.Equal(t, expected[i], v)
	}
}

func TestBroadcastToWithBroadcastingTwoDimsWithOnes(t *testing.T) {
	a := Arange(3)
	a.Reshape(&domain.IVector{Values: []int{3, 1}})
	b, err := broadcastTo(a, domain.IVector{Values: []int{3, 3, 3}}, true)
	assert.Nil(t, err)
	assert.Equal(t, 3, b.Dims)
	expected := []float64{0, 0, 0, 1, 1, 1, 2, 2, 2, 0, 0, 0, 1, 1, 1, 2, 2, 2, 0, 0, 0, 1, 1, 1, 2, 2, 2}
	for i, v := range b.Elements.Values {
		assert.Equal(t, expected[i], v)
	}
}

func TestDotBroadcastingWithTwoDims(t *testing.T) {
	// A
	//[[[0, 1],
	//  [2, 3]],
	//
	// [[4, 5],
	//  [6, 7]]]
	// Flattened: [0, 1, 2, 3, 4, 5, 6, 7]
	a := Arange(8)
	a.Reshape(&domain.IVector{Values: []int{2, 2, 2}})
	// B
	// [[0],
	//  [1]]
	// Flattened := [0,1]
	// B Broadcasted
	// [[[0, 0],
	//   [1, 1]],
	//
	//  [[0, 0],
	//   [1, 1]]]
	// Flattened :=[0, 0, 1, 1, 0, 0, 1, 1]
	b := Arange(2)
	b.Reshape(&domain.IVector{Values: []int{2, 1}})
	//C
	//  [[[0, 1],
	//    [3, 4]],
	//
	//   [[4, 5],
	//    [7, 8]]])
	c := Addition(a, b)
	assert.NotNil(t, c)
	assert.Equal(t, a.Size, c.Size)
	v := []float64{0, 1, 3, 4, 4, 5, 7, 8}
	for i := 0; i < a.Size; i++ {
		assert.Equal(t, v[i], c.Elements.Values[i])
	}
	assert.Equal(t, 2, c.Shape.Values[0])
	assert.Equal(t, 2, c.Shape.Values[1])
	assert.Equal(t, 2, c.Shape.Values[2])
}

func TestDotBroadcastingWithThreeDims(t *testing.T) {
	a := Ones("FLOAT64", 2, 2, 2)
	b := Ones("FLOAT64", 2, 1, 1)
	c := Addition(a, b)
	assert.NotNil(t, c)
	assert.Equal(t, a.Size, c.Size)
	for i := 0; i < a.Size; i++ {
		assert.Equal(t, float64(2), c.Elements.Values[i])
	}
	assert.Equal(t, 2, c.Shape.Values[0])
	assert.Equal(t, 2, c.Shape.Values[1])
	assert.Equal(t, 2, c.Shape.Values[2])
}

func TestDotBroadcastMiddleOneCase(t *testing.T) {
	a := Ones("FLOAT64", 2, 2, 2)
	b := Ones("FLOAT64", 2, 1, 2)
	c := Addition(a, b)
	assert.NotNil(t, c)
	assert.Equal(t, 8, c.Size)
}

func TestDotBroadcastWithSameDimensions(t *testing.T) {
	a := Ones("FLOAT64", 60, 2)
	b := Ones("FLOAT64", 60, 1)
	c := Addition(a, b)
	assert.NotNil(t, c)
	assert.Equal(t, 60, c.Shape.Values[0])
	assert.Equal(t, 2, c.Shape.Values[1])
}

func TestDotBroadcastWithSameDimensionsInReverse(t *testing.T) {
	a := Ones("FLOAT64", 60, 1)
	b := Ones("FLOAT64", 60, 2)
	c := Addition(a, b)
	assert.NotNil(t, c)
	assert.Equal(t, 60, c.Shape.Values[0])
	assert.Equal(t, 2, c.Shape.Values[1])
}
