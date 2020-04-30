package numpygo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Ones
func TestOnesInitializationWithFloat64(t *testing.T) {
	array := Ones("FLOAT64", 10, 10)
	assert.NotNil(t, array)
	assert.Equal(t, float64(1), array.Elements.Values[0])
	assert.Equal(t, 100, len(array.Elements.Values))
	assert.Equal(t, 100, array.Size)
	assert.Equal(t, 2, len(array.Shape.Values))
	assert.Equal(t, 10, array.Shape.Values[0])
	assert.Equal(t, 10, array.Shape.Values[1])
	assert.Equal(t, 10, array.Strides.Values[0])
	assert.Equal(t, 1, array.Strides.Values[1])
}

func TestOnesInitializationWithDefault(t *testing.T) {
	array := Ones("", 10, 10)
	assert.NotNil(t, array)
	assert.Equal(t, 0, array.Size)
}

func TestOnesInitializationWithEmptySize(t *testing.T) {
	array := Ones("FLOAT64")
	assert.NotNil(t, array)
	assert.Equal(t, 0, array.Size)
}

// zeros
func TestZerosInitializationWithFloat64(t *testing.T) {
	array := Zeros("FLOAT64", 10, 10)
	assert.NotNil(t, array)
	assert.Equal(t, float64(0), array.Elements.Values[0])
	assert.Equal(t, 100, len(array.Elements.Values))
	assert.Equal(t, 100, array.Size)
	assert.Equal(t, 2, len(array.Shape.Values))
	assert.Equal(t, 10, array.Shape.Values[0])
	assert.Equal(t, 10, array.Shape.Values[1])
}

func TestZerosInitializationWithDefault(t *testing.T) {
	array := Zeros("", 10, 10)
	assert.NotNil(t, array)
	assert.Equal(t, 0, array.Size)
}

func TestZerosInitializationWithEmptySize(t *testing.T) {
	array := Zeros("FLOAT64")
	assert.NotNil(t, array)
	assert.Equal(t, 0, array.Size)

}

// Arange
func TestArange(t *testing.T) {
	ar := Arange(0, 10, 2)
	assert.NotNil(t, ar)
	assert.Equal(t, ar.Shape.Values[0], 5)
	assert.Equal(t, float64(0), ar.Elements.Values[0])
	assert.Equal(t, float64(8), ar.Elements.Values[4])
}

func TestArangeContinuous(t *testing.T) {
	ar := Arange(10)
	assert.NotNil(t, ar)
	assert.Equal(t, 10, ar.Shape.Values[0])
	for i := 0; i < 10; i++ {
		assert.Equal(t, float64(i), ar.Elements.Values[i])
	}
}

func TestArangeWithTwoDimFail(t *testing.T) {
	ar := Arange(10, 2)
	assert.NotNil(t, ar)
	assert.Equal(t, 0, ar.Size)
}

func TestArangeWithTwoDim(t *testing.T) {
	ar := Arange(0, 10)
	assert.NotNil(t, ar)
	assert.Equal(t, 10, ar.Shape.Values[0])
	assert.Equal(t, float64(0), ar.Elements.Values[0])
	assert.Equal(t, float64(4), ar.Elements.Values[4])
}

func TestArangeWithoutSize(t *testing.T) {
	ar := Arange()
	assert.NotNil(t, ar)
	assert.Equal(t, 0, ar.Size)
}
