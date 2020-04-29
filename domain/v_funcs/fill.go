package v_funcs


func Fill(array []interface{},value interface{}) []interface{}{
	for i, _ := range array {
		array[i] = value
	}
	return array
}

func Add(array []interface{}) interface{}{
	if len(array) == 0 {
		return 0.0
	}
	value := array[0]
	for i, v := range array {
		if i!=0{
			value = value.(float64) + v.(float64)
		}
	}
	return value
}