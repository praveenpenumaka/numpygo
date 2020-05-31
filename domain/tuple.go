package domain

// Two dimensional tuple representation
// TODO: Extend this for multi dimensions
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
