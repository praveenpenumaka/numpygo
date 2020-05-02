package domain

import (
	"github.com/stretchr/testify/assert"
	"math"
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

func TestVector_Sum(t *testing.T) {
	vec := &Vector{Values: []float64{0, 1, 2, 3}}
	sum := vec.Sum()
	assert.Equal(t, 6.0, sum)
}

func TestVector_Minimum(t *testing.T) {
	vec := &Vector{Values: []float64{0, 1, 2, 3}}
	min := vec.Minimum()
	assert.Equal(t, 0.0, min)
}

func TestVector_Max(t *testing.T) {
	vec := &Vector{Values: []float64{0, 1, 2, 3}}
	max := vec.Maximum()
	assert.Equal(t, 3.0, max)
}

func TestVector_MaxWithEmptyVector(t *testing.T) {
	vec := &Vector{Values: []float64{}}
	max := vec.Maximum()
	assert.Equal(t, math.Inf(-1), max)
}

func TestVector_MinWithEmptyVector(t *testing.T) {
	vec := &Vector{Values: []float64{}}
	max := vec.Minimum()
	assert.Equal(t, math.Inf(1), max)
}
