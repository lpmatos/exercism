package greeting

// Hello is a struct to represente the hellow world message.
type Hello struct {
	message string
}

func newHello() Hello {
	return Hello{message: "Hello, World!"}
}

// HelloWorld should have a comment documenting it.
func HelloWorld() string {
	m := newHello()
	return m.message
}
