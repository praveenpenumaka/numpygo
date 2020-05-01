package domain

import (
	"github.com/praveenpenumaka/numpygo/domain/v_funcs"
	"math"
)

type Vector struct {
	Values []float64
}

func (v *Vector) Bind(vfunc func(value interface{}, args ...interface{}) interface{}, args ...interface{}) {
	for i, val := range v.Values {
		v.Values[i] = vfunc(val, args...).(float64)
	}
}

func (v *Vector) Clip(min, max float64) {
	v.Bind(v_funcs.Clip, min, max)
}

func (v *Vector) Fill(value float64) {
	v.Bind(v_funcs.Fill, value)
}

func (v *Vector) Zeros() {
	v.Fill(float64(0))
}

func (v *Vector) Add(a *Vector) float64 {
	if len(v.Values) != len(a.Values) {
		return math.Inf(1)
	}
	for i, _ := range v.Values {
		v.Values[i] = v.Values[i] + a.Values[i]
	}
	return float64(0)
}

func (v *Vector) Max() (float64, int) {
	if len(v.Values) == 0 {
		return math.Inf(1), -1
	}
	max := v.Values[0]
	maxIndex := -1
	for index, element := range v.Values {
		if element > max {
			max = element
			maxIndex = index
		}
	}
	return max, maxIndex
}

func (v *Vector) Min() (float64, int) {
	if len(v.Values) == 0 {
		return math.Inf(-1), -1
	}
	min := v.Values[0]
	minIndex := -1
	for index, element := range v.Values {
		if element > min {
			min = element
			minIndex = index
		}
	}
	return min, minIndex
}

func (v *Vector) Unique() *Vector {
	var uniqueMap map[float64]bool
	uniqueMap = make(map[float64]bool)
	var keys []float64
	for _, e := range v.Values {
		uniqueMap[e] = true
	}
	for key := range uniqueMap {
		keys = append(keys, key)
	}
	return &Vector{Values: keys}
}

func (v *Vector) Pow(p float64) {
	for i, element := range v.Values {
		v.Values[i] = math.Pow(element, p)
	}
}

func (v *Vector) Sum() float64 {
	sum := float64(0)
	for _, element := range v.Values {
		sum += element
	}
	return sum
}

func (v *Vector) Mean() float64 {
	sum := float64(0)
	count := 0
	for _, element := range v.Values {
		sum += element
		count++
	}
	if count == 0 {
		return float64(0)
	}
	return sum / float64(count)
}
