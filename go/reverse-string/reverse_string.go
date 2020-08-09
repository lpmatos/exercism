package reverse

// Message struct.
type Message struct {
	content string
}

// Reverse function that reverse a string.
func Reverse(content string) (result string) {
	message := Message{content: content}
	for _, value := range message.content {
		result = string(value) + result
	}
	return
}
