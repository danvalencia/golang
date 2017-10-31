package ds

// Uniq receives a slice and returns a new slice without any duplicates
func Uniq(arr []interface{}) []interface{} {
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
