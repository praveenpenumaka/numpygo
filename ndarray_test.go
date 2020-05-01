package numpygo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClip(t *testing.T) {
	a := Ones("FLOAT64", 10, 3)
	b := Clip(a, 3, 4)
	assert.NotNil(t, b)
	assert.Equal(t, 30, b.Size)
	assert.Equal(t, 10, b.Shape.Values[0])
	assert.Equal(t, 3, b.Shape.Values[1])
	for _, v := range b.Elements.Values {
		assert.Equal(t, 3.0, v)
	}
}
