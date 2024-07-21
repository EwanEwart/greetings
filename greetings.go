package greetings

import (
	"fmt"
	// "golang.org/x/text/message"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi %v, at Go labs.", name)
	return message
}
