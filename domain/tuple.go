package domain

type Tuple struct {
	X int
	Y int
}

func (tuple *Tuple) IsValidDimension() bool {
	if tuple.X < 0 {
		return false
	}
	if tuple.Y < 0 {
		return false
	}
	if tuple.X > tuple.Y {
		return false
	}
	return true
}

func (tuple *Tuple) IsValidDimensionWithSize(size int) bool {
	return tuple.IsValidDimension() && tuple.Y >= size
}
