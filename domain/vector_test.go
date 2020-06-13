package domain

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestVector_FillIndex(t *testing.T) {
	vec := &Vector{Values: []float64{0, 1, 2, 3}}
	vec.FillIndex([]float64{5.0}, 1)
	assert.Equal(t, 5.0, vec.Values[1])
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

func TestVector_Maximum(t *testing.T) {
	vec := &Vector{Values: []float64{0, 1, 2, 3}}
	max := vec.Maximum()
	assert.Equal(t, 3.0, max)
}

func TestVector_MaximumWithEmptyVector(t *testing.T) {
	vec := &Vector{Values: []float64{}}
	max := vec.Maximum()
	assert.Equal(t, math.Inf(-1), max)
}

func TestVector_MinimumWithEmptyVector(t *testing.T) {
	vec := &Vector{Values: []float64{}}
	max := vec.Minimum()
	assert.Equal(t, math.Inf(1), max)
}

func TestVector_MinWithEmptyVector(t *testing.T) {
	vec := &Vector{Values: nil}
	v, ind := vec.Min()
	assert.Equal(t, math.Inf(-1), v)
	assert.Equal(t, -1, ind)
}

func TestVector_Min(t *testing.T) {
	vec := &Vector{Values: []float64{50, 10, 20, 30}}
	v, ind := vec.Min()
	assert.Equal(t, 10.0, v)
	assert.Equal(t, 1, ind)
}

func TestVector_MaxWithEmptyVector(t *testing.T) {
	vec := &Vector{Values: nil}
	v, ind := vec.Max()
	assert.Equal(t, math.Inf(1), v)
	assert.Equal(t, -1, ind)
}

func TestVector_Max(t *testing.T) {
	vec := &Vector{Values: []float64{10, 20, 30}}
	v, ind := vec.Max()
	assert.Equal(t, 30.0, v)
	assert.Equal(t, 2, ind)
}

func TestVector_UniqueWithEmptyVector(t *testing.T) {
	vec := &Vector{}
	v := vec.Unique()
	assert.NotNil(t, v)
	assert.Equal(t, 0, len(v.Values))
}

func TestVector_Unique(t *testing.T) {
	vec := &Vector{Values: []float64{10, 10, 20, 20, 30}}
	v := vec.Unique()
	assert.NotNil(t, v)
	assert.Equal(t, 3, len(v.Values))
	assert.Equal(t, 60.0, v.Sum())
}

func TestVector_MeanWithEmpty(t *testing.T) {
	vec := &Vector{}
	mean := vec.Mean()
	assert.Equal(t, 0.0, mean)
}

func TestVector_Mean(t *testing.T) {
	vec := &Vector{[]float64{10, 20, 30}}
	mean := vec.Mean()
	assert.Equal(t, 20.0, mean)
}

func TestVector_exp(t *testing.T) {
	vec := &Vector{[]float64{10, 10, 10}}
	vec.Exp()
	assert.Equal(t, math.Exp(10), vec.Values[0])
}

func TestVector_exp2(t *testing.T) {
	vec := &Vector{[]float64{10, 10, 10}}
	vec.Exp2()
	assert.Equal(t, math.Exp2(10), vec.Values[0])
}

func TestVector_log(t *testing.T) {
	vec := &Vector{[]float64{10, 10, 10}}
	vec.Log()
	assert.Equal(t, math.Log(10), vec.Values[0])
}

func TestVector_log2(t *testing.T) {
	vec := &Vector{[]float64{10, 10, 10}}
	vec.Log2()
	assert.Equal(t, math.Log2(10), vec.Values[0])
}

func TestVector_pow(t *testing.T) {
	vec := &Vector{[]float64{10, 10, 10}}
	vec.Pow(2.0)
	assert.Equal(t, math.Pow(10, 2), vec.Values[0])
}
