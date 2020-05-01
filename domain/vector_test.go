package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVector_Clip(t *testing.T) {
	vec := &Vector{Values: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	vec.Clip(3.0, 8.0)
	for _, v := range vec.Values {
		assert.True(t, v >= 3.0)
		assert.True(t, v <= 8.0)
	}
}

func TestVector_Fill(t *testing.T) {
	vec := &Vector{Values: []float64{0, 1, 2, 3}}
	vec.Fill(3.0)
	for _, v := range vec.Values {
		assert.Equal(t, 3.0, v)
	}
}
