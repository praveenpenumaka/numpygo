package domain

type IVector struct {
	Values []int
}

func (v *IVector) CopyFrom(c []int) {
	if c == nil || len(c) == 0 {
		return
	}
	if v.Values == nil {
		v.Values = make([]int, len(c))
	}
	for i := range v.Values {
		v.Values[i] = c[i]
	}
}

func (v *IVector) Zeros() {
	for i, _ := range v.Values {
		v.Values[i] = 0
	}
}

// Returns a new vector removing element at i
func (v *IVector) Remove(index int) *IVector {
	var newVector IVector
	for i, _ := range v.Values {
		if i != index {
			newVector.Values = append(newVector.Values, v.Values[i])
		}
	}
	return &newVector
}

func (v *IVector) Add(a *IVector) {
	if len(v.Values) != len(a.Values) {
		return
	}
	for i, _ := range v.Values {
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

func (v *IVector) Min() int {
	if len(v.Values) == 0 {
		return 0
	}
	minValue := v.Values[0]
	for _, element := range v.Values {
		if element < minValue {
			minValue = element
		}
	}
	return minValue
}

func (v *IVector) Equals(a *IVector) bool {
	if len(v.Values) != len(a.Values) {
		return false
	}
	for i, _ := range v.Values {
		if v.Values[i] != a.Values[i] {
			return false
		}
	}
	return true
}
