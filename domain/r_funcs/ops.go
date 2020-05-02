package r_funcs

func Sum(red, value interface{}, args ...interface{}) interface{} {
	return red.(float64) + value.(float64)
}

func Min(red, value interface{}, args ...interface{}) interface{} {
	if value.(float64) < red.(float64) {
		return value
	}
	return red
}

func Max(red, value interface{}, args ...interface{}) interface{} {
	if value.(float64) > red.(float64) {
		return value
	}
	return red
}
