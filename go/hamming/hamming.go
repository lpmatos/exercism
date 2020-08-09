package hamming

import "errors"

// Distance function - Unicode is a Rune in Golang.
func Distance(a, b string) (int, error) {
	lens, difference := len(a), 0
	if lens != len(b) {
		return -1, errors.New("Error - different lenghts")
	}
	for index := 0; index < lens; index++ {
		if a[index] != b[index] {
			difference++
		}
	}
	return difference, nil
}
