package domain

import (
	"github.com/praveenpenumaka/numpygo/domain/r_funcs"
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

func (v *Vector) Reduce(rfunc func(red, value interface{}, args ...interface{}) interface{}, args ...interface{}) float64 {
	reduced := args[0].(float64)
	for _, val := range v.Values {
		reduced = rfunc(reduced, val, args...).(float64)
	}
	return reduced
}

func (v *Vector) Clip(min, max float64) {
	v.Bind(v_funcs.Clip, min, max)
}

func (v *Vector) Fill(value float64) {
	v.Bind(v_funcs.Fill, value)
}

func (v *Vector) FillIndex(values []float64, index int) {
	vLen := len(v.Values)
	for i, val := range values {
		rIndex := index + i
		if rIndex < vLen {
			v.Values[index+i] = val
		}
	}
}

func (v *Vector) Zeros() {
	v.Bind(v_funcs.Fill, 0.0)
}

func (v *Vector) Maximum() float64 {
	return v.Reduce(r_funcs.Max, math.Inf(-1))
}

func (v *Vector) Minimum() float64 {
	return v.Reduce(r_funcs.Min, math.Inf(1))
}

func (v *Vector) Exp() {
	v.Bind(v_funcs.Exp)
}

func (v *Vector) Exp2() {
	v.Bind(v_funcs.Exp2)
}

func (v *Vector) Log() {
	v.Bind(v_funcs.Log)
}

func (v *Vector) Log2() {
	v.Bind(v_funcs.Log2)
}

func (v *Vector) Pow(p float64) {
	v.Bind(v_funcs.Pow, p)
}

func (v *Vector) Sum() float64 {
	return v.Reduce(r_funcs.Sum, float64(0))
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
	maxIndex := 0
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
	minIndex := 0
	for index, element := range v.Values {
		if element < min {
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
