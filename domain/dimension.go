package domain

type Dimensions struct {
	Dims []Tuple
}

func (d *Dimensions) IsValid() bool {
	if d.Dims == nil || len(d.Dims) == 0 {
		return false
	}
	isValid := true
	for _, dim := range d.Dims {
		isValid = isValid && dim.IsValidDimension()
	}
	return isValid
}

func (d *Dimensions) GetTotalSize() int {
	totalSize := 0
	for _, dim := range d.Dims {
		if !dim.IsValidDimension() {
			return -1
		}
		dimSize := (dim.Y - dim.X) + 1
		totalSize += dimSize
	}
	return totalSize
}

func (d *Dimensions) GetShape() *IVector {
	var shape IVector
	for _, dim := range d.Dims {
		if !dim.IsValidDimension() {
			return nil
		}
		dimSize := (dim.Y - dim.X) + 1
		shape.Values = append(shape.Values, dimSize)
	}
	return &shape
}

func (d *Dimensions) GetStartVector() *IVector {
	var startVector IVector
	for _, dim := range d.Dims {
		if !dim.IsValidDimension() {
			return nil
		}
		startVector.Values = append(startVector.Values, dim.X)
	}
	return &startVector
}
