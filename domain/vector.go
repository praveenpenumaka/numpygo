package domain

import "math"

type Vector struct {
	Values []float64
}

func (v *Vector) Fill(value float64) {
	for i, _ := range v.Values {
		v.Values[i] = value
	}
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

func (v *Vector) Max() float64 {
	if len(v.Values) == 0{
		return math.Inf(1)
	}
	max := v.Values[0]
	for _, element := range v.Values {
		if element > max {
			max = element
		}
	}
	return max
}

func (v *Vector) Min() float64 {
	if len(v.Values) == 0{
		return math.Inf(-1)
	}
	min := v.Values[0]
	for _, element := range v.Values {
		if element > min {
			min = element
		}
	}
	return min
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

func (v *Vector) Mean() float64  {
	sum := float64(0)
	count :=0
	for _, element := range v.Values {
		sum += element
		count++
	}
	if count==0 {
		return float64(0)
	}
	return sum/float64(count)
}
