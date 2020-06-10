package domain

type IVector struct {
	Values []int
}

func (v *IVector) CopyFrom(c []int) bool {
	if len(c) == 0 {
		return false
	}
	if v.Values == nil {
		v.Values = make([]int, len(c))
	} else if len(c) != len(v.Values) {
		return false
	}
	for i := range v.Values {
		v.Values[i] = c[i]
	}
	return true
}

func (v *IVector) Padding(location string, value, len int) {
	if len <= 0 || location == "" {
		return
	}
	d := make([]int, len)
	for i := range d {
		d[i] = value
	}
	if location == "LEAD" {
		v.Values = append(d, v.Values...)
	} else if location == "TRAIL" {
		v.Values = append(v.Values, d...)
	}
}

func (v *IVector) Zeros() {
	for i := range v.Values {
		v.Values[i] = 0
	}
}

// Returns a new vector removing element at i
func (v *IVector) Remove(index int) *IVector {
	if index < 0 || index >= len(v.Values) {
		return nil
	}
	values := v.Values
	newVector := &IVector{}
	indexN := index + 1
	newVector.Values = append(newVector.Values, values[:index]...)
	newVector.Values = append(newVector.Values, values[indexN:]...)
	return newVector
}

func (v *IVector) Add(a *IVector) {
	if len(v.Values) == 0 || len(a.Values) == 0 || len(v.Values) != len(a.Values) {
		return
	}
	for i := range v.Values {
		v.Values[i] = v.Values[i] + a.Values[i]
	}
}

func (v *IVector) Mult() int {
	multValue := 1
	for _, element := range v.Values {
		multValue = multValue * element
	}
	return multValue
}

func (v *IVector) Min() (int, int) {
	if len(v.Values) == 0 {
		return 0, -1
	}
	minIndex := 0
	minValue := v.Values[minIndex]
	for index, element := range v.Values {
		if element < minValue {
			minValue = element
			minIndex = index
		}
	}
	return minValue, minIndex
}

func (v *IVector) Max() (int, int) {
	if len(v.Values) == 0 {
		return 0, -1
	}
	maxIndex := 0
	maxValue := v.Values[maxIndex]
	for index, element := range v.Values {
		if element > maxValue {
			maxValue = element
			maxIndex = index
		}
	}
	return maxValue, maxIndex
}

func (v *IVector) Equals(a *IVector) bool {
	if len(v.Values) != len(a.Values) {
		return false
	}
	for i := range v.Values {
		if v.Values[i] != a.Values[i] {
			return false
		}
	}
	return true
}
