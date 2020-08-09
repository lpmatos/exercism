// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// Valor de um gigasecond.
const gigasecond = 1e9 * time.Second

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	// Dado o momento 't', adicionamos um 'gigasecond' para determinar quando ele aconteceria após o gigasecond ter passado.
	// A função Add adiciona um tempo a outro tempo.
	return t.Add(gigasecond)
}
