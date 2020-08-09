package twofer

import "fmt"

// Person representation.
type Person struct {
	name string
}

func check(value string) string {
	if value == "" {
		return "you"
	} else {
		return value
	}
}

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	return fmt.Sprintf("One for %s, one for me.", Person{name: check(name)}.name)
}
