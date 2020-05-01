package v_funcs

func Clip(v interface{}, args ...interface{}) interface{} {
	if len(args) != 2 {
		panic("Invalid number of arguments for v_func:clip")
	}
	if v.(float64) < args[0].(float64) {
		return args[0]
	}
	if v.(float64) > args[1].(float64) {
		return args[1]
	}
	return v
}
