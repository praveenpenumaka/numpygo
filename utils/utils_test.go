package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVector(t *testing.T) {
	vec1, err1 := GetVector(0, []int{4, 2, 1}, []int{2, 2, 2})
	vec2, err2 := GetVector(1, []int{4, 2, 1}, []int{2, 2, 2})
	vec3, err3 := GetVector(3, []int{4, 2, 1}, []int{2, 2, 2})
	assert.NotNil(t, vec1)
	assert.Nil(t, err1)
	assert.NotNil(t, vec2)
	assert.Nil(t, err2)
	assert.NotNil(t, vec3)
	assert.Nil(t, err3)
	assert.Equal(t, 3, len(vec1))
	assert.Equal(t, 3, len(vec2))
	assert.Equal(t, 3, len(vec3))
	i1, _ := GetIndex(vec1, []int{4, 2, 1}, []int{2, 2, 2})
	assert.Equal(t, 0, i1)
	i2, _ := GetIndex(vec2, []int{4, 2, 1}, []int{2, 2, 2})
	assert.Equal(t, 1, i2)
	i3, _ := GetIndex(vec3, []int{4, 2, 1}, []int{2, 2, 2})
	assert.Equal(t, 3, i3)
}

func TestGetVectorWithInvalidIndex(t *testing.T) {
	_, err := GetVector(9, []int{4, 2, 1}, []int{2, 2, 2})
	assert.NotNil(t, err)
}

func TestGetIndex(t *testing.T) {
	index1, err1 := GetIndex([]int{0, 0, 0}, []int{4, 2, 1}, []int{2, 2, 2})
	index2, err2 := GetIndex([]int{0, 0, 1}, []int{4, 2, 1}, []int{2, 2, 2})
	_, err3 := GetIndex([]int{0, 0, 2}, []int{4, 2, 1}, []int{2, 2, 2})
	assert.Nil(t, err1)
	assert.Equal(t, 0, index1)
	assert.Nil(t, err2)
	assert.Equal(t, 1, index2)
	assert.NotNil(t, err3)
}
