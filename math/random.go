package math

import (
	"math/rand"
	"time"

	"github.com/universe-sh/utils-go/slice"
)

// RandInt new integer
func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

// RandIntArray new random integer array
func RandIntArray(input []int, n int) []int {
	var result []int

	if n >= len(input) {
		return input
	} else if len(input) <= 0 {
		return []int{}
	}

	for i := 0; i < n; {
		j := RandInt(0, 9)
		if !slice.IntegerInSlice(input[j], result) {
			result = append(result, input[j])
			i++
		}
	}

	return result
}
