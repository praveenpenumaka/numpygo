package numpygo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNDIndex(t *testing.T) {
	ndIndex := NewNDIndex([]int{2, 2, 2})
	assert.NotNil(t, ndIndex)
	ndIndex.Reset()
	for i := 0; i < 3; i++ {
		newIndex := ndIndex.Next()
		assert.Equal(t, 3, len(newIndex.Values))
		assert.Equal(t, 0, newIndex.Values[0])
		if i < 2 {
			assert.Equal(t, 0, newIndex.Values[1])
		} else {
			assert.Equal(t, 1, newIndex.Values[1])
		}
		if i < 2 {
			assert.Equal(t, i, newIndex.Values[2])
		} else {
			assert.Equal(t, 0, newIndex.Values[2])
		}
	}
}

func TestCreateNDIndexForOneD(t *testing.T){
	ndIndex := NewNDIndex([]int{2, 1})
	assert.NotNil(t, ndIndex)
	ndIndex.Reset()
	v:=ndIndex.Next()
	assert.Equal(t,0,v.Values[0])
	assert.Equal(t,0,v.Values[1])
	v=ndIndex.Next()
	assert.Equal(t,1,v.Values[0])
	assert.Equal(t,0,v.Values[1])
	v=ndIndex.Next()
	assert.Nil(t,v)
}

func TestCreateNDIndexForOneDTrns(t *testing.T){
	ndIndex := NewNDIndex([]int{1, 2})
	assert.NotNil(t, ndIndex)
	ndIndex.Reset()
	v:=ndIndex.Next()
	assert.Equal(t,0,v.Values[0])
	assert.Equal(t,0,v.Values[1])
	v=ndIndex.Next()
	assert.Equal(t,0,v.Values[0])
	assert.Equal(t,1,v.Values[1])
	v=ndIndex.Next()
	assert.Nil(t,v)
}