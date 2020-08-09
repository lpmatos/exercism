package bob

import (
	"strings"
	"unicode"
)

/*
Bob is a lackadaisical teenager. In conversation, his responses are very limited.

Bob answers 'Sure.' if you ask him a question, such as "How are you?".

He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).

He answers 'Calm down, I know what I'm doing!' if you yell a question at him.

He says 'Fine. Be that way!' if you address him without actually saying anything.

He answers 'Whatever.' to anything else.

Bob's conversational partner is a purist when it comes to written communication and always follows normal rules regarding sentence punctuation in English.
*/

// Message is a convenience type for identifying the message
type Message struct {
	message string
}

func (m Message) isEmpty() bool {
	return m.message == ""
}

func (m Message) isUpcased() bool {
	return strings.ToUpper(m.message) == m.message
}

func (m Message) onlyLetters() bool {
	return strings.IndexFunc(m.message, unicode.IsLetter) >= 0
}

func (m Message) isQuestion() bool {
	return strings.HasSuffix(m.message, "?")
}

func (m Message) isScreaming() bool {
	return m.onlyLetters() && m.isUpcased()
}

func (m Message) isCrazy() bool {
	return m.isScreaming() && m.isQuestion()
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	message := Message{
		message: strings.TrimSpace(remark),
	}

	switch {
	case message.isEmpty():
		return "Fine. Be that way!"
	case message.isCrazy():
		return "Calm down, I know what I'm doing!"
	case message.isScreaming():
		return "Whoa, chill out!"
	case message.isQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}
