package greetings

//  error : the Go standard library's errors package
import (
	"errors"
	"fmt"
	// "golang.org/x/text/message"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi %v, at Go labs.", name)
	return message, nil
}
