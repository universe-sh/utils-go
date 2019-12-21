package maps

// IndexExist in map
func IndexExist(data map[string]string, index string) bool {
	if _, ok := data[index]; ok {
		return true
	}

	return false
}
