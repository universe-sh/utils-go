package slice

// StringInSlice func
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

// MapCounterInSlice func
func MapCounterInSlice(a map[string]int64, list []string) bool {
	if len(a) > len(list) {
		return false
	}

	for k, v := range a {
		for _, b := range list {
			if b == k {
				v--
			}
		}

		if v > 0 {
			return false
		}
	}

	return true
}
