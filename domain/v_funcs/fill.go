package v_funcs

func Fill(value interface{}, args ...interface{}) interface{} {
	return args[0]
}

func Add(value interface{}, args ...interface{}) interface{} {
	if len(args) == 0 {
		return value
	}
	for i, _ := range args {
		value = value.(float64) + args[i].(float64)
	}
	return value
}
