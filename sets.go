package main

func uniq(arr []interface{}) []interface{} {
	uniqueMap := make(map[interface{}]bool, len(arr))
	uniqueSlice := make([]interface{}, 0)

	for _, ele := range arr {
		if !uniqueMap[ele] {
			uniqueSlice = append(uniqueSlice, ele)
			uniqueMap[ele] = true
		}
	}

	return uniqueSlice
}
