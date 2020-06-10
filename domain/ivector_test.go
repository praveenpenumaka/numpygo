package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorPaddingWithInvalidIdentifier(t *testing.T) {
	v := &IVector{Values: []int{10}}
	v.Padding("", 10, 10)
	assert.Equal(t, 1, len(v.Values))
}

func TestVectorPaddingWithInvalidLength(t *testing.T) {
	v := &IVector{Values: []int{10}}
	v.Padding("", 10, -1)
	assert.Equal(t, 1, len(v.Values))
}

func TestVectorPaddingWithLeads(t *testing.T) {
	v := &IVector{Values: []int{20, 30}}
	v.Padding("LEAD", 10, 2)
	assert.Equal(t, 4, len(v.Values))
	assert.Equal(t, 10, v.Values[0])
	assert.Equal(t, 10, v.Values[1])
	assert.Equal(t, 20, v.Values[2])
	assert.Equal(t, 30, v.Values[3])
}

func TestVectorPaddingWithTrail(t *testing.T) {
	v := &IVector{Values: []int{20, 30}}
	v.Padding("TRAIL", 40, 2)
	assert.Equal(t, 4, len(v.Values))
	assert.Equal(t, 20, v.Values[0])
	assert.Equal(t, 30, v.Values[1])
	assert.Equal(t, 40, v.Values[2])
	assert.Equal(t, 40, v.Values[3])
}

func TestVector_Zeros(t *testing.T) {
	v := &IVector{Values: []int{10, 10, 20}}
	v.Zeros()
	assert.Equal(t, 3, len(v.Values))
	for _, v := range v.Values {
		assert.Equal(t, 0, v)
	}
}

func TestVector_CopyFromWithZeroLen(t *testing.T) {
	v := &IVector{Values: []int{10, 20, 30}}
	v.CopyFrom(nil)
	assert.Equal(t, 3, len(v.Values))
}

func TestVector_CopyFromWithInvalidLen(t *testing.T) {
	v := &IVector{Values: []int{10, 20, 30}}
	v.CopyFrom([]int{20})
	assert.Equal(t, 3, len(v.Values))
}

func TestVector_CopyFromSuccess(t *testing.T) {
	v := &IVector{Values: []int{10, 20, 30}}
	v.CopyFrom([]int{2, 3, 4})
	assert.Equal(t, 3, len(v.Values))
	assert.Equal(t, 2, v.Values[0])
	assert.Equal(t, 3, v.Values[1])
	assert.Equal(t, 4, v.Values[2])
}

func TestVector_CopyFromWithInitialize(t *testing.T) {
	v := &IVector{Values: nil}
	v.CopyFrom([]int{2, 3, 4})
	assert.Equal(t, 3, len(v.Values))
	assert.Equal(t, 2, v.Values[0])
	assert.Equal(t, 3, v.Values[1])
	assert.Equal(t, 4, v.Values[2])
}

func TestVector_RemoveWithInvalidVector(t *testing.T) {
	v := &IVector{Values: nil}
	x := v.Remove(2)
	assert.Nil(t, x)
}

func TestVector_RemoveWithInvalidIndex(t *testing.T) {
	v := &IVector{Values: []int{10, 20, 30}}
	x := v.Remove(5)
	assert.Nil(t, x)
}

func TestVector_RemoveWithValidIndex(t *testing.T) {
	v := &IVector{Values: []int{10, 20, 30}}
	x1 := v.Remove(0)
	x2 := v.Remove(1)
	x3 := v.Remove(2)
	assert.NotNil(t, x1)
	assert.NotNil(t, x2)
	assert.NotNil(t, x3)
	assert.Equal(t, 2, len(x1.Values))
	assert.Equal(t, 2, len(x2.Values))
	assert.Equal(t, 2, len(x3.Values))
	assert.Equal(t, 20, x1.Values[0])
	assert.Equal(t, 30, x1.Values[1])

	assert.Equal(t, 10, x2.Values[0])
	assert.Equal(t, 30, x2.Values[1])

	assert.Equal(t, 10, x3.Values[0])
	assert.Equal(t, 20, x3.Values[1])
}

func TestVector_AddWithInvalidVector(t *testing.T) {
	v := &IVector{Values: []int{10, 20, 30}}
	a := &IVector{Values: nil}
	v.Add(a)
	assert.Equal(t, 3, len(v.Values))
	assert.Equal(t, 10, v.Values[0])
	assert.Equal(t, 20, v.Values[1])
	assert.Equal(t, 30, v.Values[2])
}

func TestVector_AddWithValidVector(t *testing.T) {
	v := &IVector{Values: []int{10, 20, 30}}
	a := &IVector{Values: []int{30, 20, 10}}
	v.Add(a)
	assert.Equal(t, 3, len(v.Values))
	assert.Equal(t, 40, v.Values[0])
	assert.Equal(t, 40, v.Values[1])
	assert.Equal(t, 40, v.Values[2])
}

func TestVector_MinIWithEmpty(t *testing.T) {
	v := &IVector{Values: []int{}}
	min, ind := v.Min()
	assert.Equal(t, 0, min)
	assert.Equal(t, -1, ind)
}

func TestVector_MinI(t *testing.T) {
	v := &IVector{Values: []int{50, 10, 20, 30}}
	min, ind := v.Min()
	assert.Equal(t, 10, min)
	assert.Equal(t, 1, ind)
}

func TestVector_MaxIWithEmpty(t *testing.T) {
	v := &IVector{Values: []int{}}
	min, ind := v.Max()
	assert.Equal(t, 0, min)
	assert.Equal(t, -1, ind)
}

func TestVector_MaxI(t *testing.T) {
	v := &IVector{Values: []int{10, 20, 30}}
	min, ind := v.Max()
	assert.Equal(t, 30, min)
	assert.Equal(t, 2, ind)
}

func TestIVector_EqualsWithInequalVectors(t *testing.T) {
	a := &IVector{Values: []int{10, 20, 30}}
	b := &IVector{Values: []int{10, 20, 30, 40}}
	assert.False(t, b.Equals(a))
}

func TestIVector_EqualsWithUnEqualVector(t *testing.T) {
	a := &IVector{Values: []int{10, 20, 30}}
	b := &IVector{Values: []int{20, 20, 20}}
	assert.False(t, b.Equals(a))
}

func TestIVector_EqualsWithEqualVector(t *testing.T) {
	a := &IVector{Values: []int{10, 20, 30}}
	b := &IVector{Values: []int{10, 20, 30}}
	assert.True(t, b.Equals(a))
}
