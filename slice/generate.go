package slice

// GenerateCombinations func
func GenerateCombinations(iter []string, length int) <-chan []string {
	c := make(chan []string)

	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan []string) {
		defer close(c)
		addIter(c, []string{}, iter, length)
	}(c)

	return c
}

// AddIter adds a letter to the combination to create a new combination.
// This new combination is passed on to the channel before we call AddIter once again
func addIter(c chan []string, combo []string, iter []string, length int) {
	// Check if we reached the length limit
	// If so, we just return without adding anything
	if length <= 0 {
		return
	}

	var newCombo []string
	for _, ch := range iter {
		newCombo = append(combo, string(ch))
		c <- newCombo
		addIter(c, newCombo, iter, length-1)
	}
}
