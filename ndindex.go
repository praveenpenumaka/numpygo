package numpygo

import "github.com/praveenpenumaka/numpygo/domain"

type NDIndex struct {
	iterationStarted bool
	currentVector    *domain.IVector
	shape            *domain.IVector
}

func NewNDIndex(shape []int) *NDIndex {
	ndIndex := &NDIndex{
		iterationStarted: false,
		shape:            &domain.IVector{Values: nil},
		currentVector:    &domain.IVector{Values: nil},
	}
	ndIndex.shape.CopyFrom(shape)
	ndIndex.currentVector.CopyFrom(shape)
	ndIndex.Reset()
	return ndIndex
}

func (ndIndex *NDIndex) Reset() {
	ndIndex.currentVector.Zeros()
}

func (ndIndex *NDIndex) IncrementIndex() bool {
	l := len(ndIndex.shape.Values)
	for i := l - 1; i >= 0; i-- {
		if ndIndex.currentVector.Values[i]+1 == ndIndex.shape.Values[i] {
			ndIndex.currentVector.Values[i] = 0
			if i == 0 {
				return false
			}
		} else {
			ndIndex.currentVector.Values[i]++
			return true
		}
	}
	return true
}

// This returns new object
// So that application can use the newly created object
func (ndIndex *NDIndex) Next() *domain.IVector {
	if !ndIndex.iterationStarted {
		ndIndex.iterationStarted = true
		nv := &domain.IVector{Values: nil}
		nv.CopyFrom(ndIndex.currentVector.Values)
		return nv
	}
	if ndIndex.IncrementIndex() {
		nv := &domain.IVector{Values: nil}
		nv.CopyFrom(ndIndex.currentVector.Values)
		return nv
	}
	return nil
}
