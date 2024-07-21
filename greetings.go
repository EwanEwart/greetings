package greetings

//  errors package : the Go standard library's errors package
//	math/rand package : to generate a random number for selecting an item from the slice.
import (
	"errors"
	"fmt"
	"math/rand"
	// "golang.org/x/text/message"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	/*
		In Hello, call the randomFormat function
		to get a format for the message you'll return,
		then use the 'format' and 'name' value together
		to create the message:
		Return a greeting that embeds the name in a message.
	*/
	message := fmt.Sprintf(randomFormat(), name)
	/*
		Return the message, or an error, as before.
	*/
	return message, nil
}

/*
Add a randomFormat function
that returns a randomly selected format as a greeting message.
Note that randomFormat starts with "a lowercase letter",
making it accessible only to code in its own package,
in other words, it's not exported.
randomFormat returns one of a set of greeting messages.
The returned message is selected at random.
*/
func randomFormat() string {
	/*
		In randomFormat,
		declare a formats "slice" with three message formats.
		When declaring a slice,
		you omit its size in the brackets,
		like this: []string.
		This tells Go that the size of the array
		underlying the slice should be dynamically changeable.
		A slice of message formats:
	*/
	formats := []string{
		"Greetings, %v, & welcome!",
		"Great to meet you again, %v!",
		"Have a nice day, %v! Wish you success!",
	}
	/* Return a randomly selected message format
	by specifying a random index
	for the slice of formats.
	*/
	return formats[rand.Intn(len(formats))]
}
