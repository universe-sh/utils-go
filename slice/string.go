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

// Combrep func
// thanks https://rosettacode.org/wiki/Combinations_with_repetitions#Go
func Combrep(lst []string, n int) [][]string {
	if n == 0 {
		return [][]string{nil}
	}
	if len(lst) == 0 {
		return nil
	}
	r := Combrep(lst[1:], n)
	for _, x := range Combrep(lst, n-1) {
		r = append(r, append(x, lst[0]))
	}
	return r
}
