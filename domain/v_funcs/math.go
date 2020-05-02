package v_funcs

import "math"

func Exp(val interface{}, args ...interface{}) interface{} {
	return math.Exp(val.(float64))
}

func Exp2(val interface{}, args ...interface{}) interface{} {
	return math.Exp2(val.(float64))
}

func Log2(val interface{}, args ...interface{}) interface{} {
	return math.Log2(val.(float64))
}

func Log(val interface{}, args ...interface{}) interface{} {
	return math.Log(val.(float64))
}

func Add(val interface{}, args ...interface{}) interface{} {
	finalValue := val.(float64)
	for _, v := range args {
		finalValue = finalValue + v.(float64)
	}
	return finalValue
}

func Sub(val interface{}, args ...interface{}) interface{} {
	finalValue := val.(float64)
	for _, v := range args {
		finalValue = finalValue - v.(float64)
	}
	return finalValue
}

func Mult(val interface{}, args ...interface{}) interface{} {
	finalValue := val.(float64)
	for _, v := range args {
		finalValue = finalValue * v.(float64)
	}
	return finalValue
}

func Div(val interface{}, args ...interface{}) interface{} {
	return val.(float64) / args[0].(float64)
}

func Pow(val interface{}, args ...interface{}) interface{} {
	return math.Pow(val.(float64), args[0].(float64))
}
