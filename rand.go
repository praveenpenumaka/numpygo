package numpygo

import (
	"github.com/praveenpenumaka/numpygo/domain"
	"math/rand"
)

const OneBigNumber = 1000000

func getRandomInteger(max int) int {
	bigRandomFloatValue := rand.Float64() * OneBigNumber
	return int(bigRandomFloatValue) % max
}

func Rand(dtype string, shape ...int) NDArray {
	nd := newNDArray(dtype, domain.IVector{Values: shape})
	if nd.Size == 0 {
		return NDArray{}
	}
	for index, _ := range nd.Elements.Values {
		nd.Elements.Values[index] = rand.Float64()
	}
	return nd
}

//This function only shuffles the array along the first axis of a multi-dimensional array.
func Shuffle(nd NDArray) NDArray {

	indexes := Arange(nd.Shape.Values[0])

	for i := 0; i < indexes.Size; i++ {
		// pick one random value between 0 and size
		randomIndex := getRandomInteger(indexes.Size)
		// Replace elementAt[i] with elementAt[rand]
		temp := indexes.Elements.Values[i]
		indexes.Elements.Values[i] = indexes.Elements.Values[randomIndex]
		indexes.Elements.Values[randomIndex] = temp
	}

	return nd.GetIndexed(indexes)
}
