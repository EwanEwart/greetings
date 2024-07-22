package greetings

//  errors package : the Go standard library's errors package
//	math/rand package : to generate a random number for selecting an item from the slice.
import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
Add a Hellos function
whose parameter is a slice of names rather than a single name.
Also, you change one of its return types
from a string to a map
so you can return names mapped to greeting messages.

Hellos returns a map
that associates each of the named people
with a greeting message,
and errors
*/
func Hellos(names []string) (map[string]string, error) {
	/*
		Create a messages map
		to associate each of the received names (as a key)
		with a generated message (as a value).

		In Go, you "initialise" a map with the following syntax:
		make(map[key-type]value-type).

		You have the Hellos function return this map to the caller.

		For more about maps,
		see Go maps in action on the Go blog:
		https://go.dev/blog/maps

		A map to associate names with messages.
	*/
	messages := make(map[string]string)
	/*
		Loop through the names your function received,
		checking that each has a non-empty value,
		then associate a message with each.

		In this for loop,
		range returns two values:
		the index of the current item in the loop
		and a copy of the item's value.

		You don't need the index,
		so you use the Go "blank identifier"
		(an underscore) to ignore it.
		For more, see The blank identifier in Effective Go:
		https://go.dev/doc/effective_go.html#blank

		Loop through the received slice of names,
		calling the Hello function to get a message for each name.
	*/
	for _, name := range names {
		/*
			Have the new Hellos function
			call the existing Hello function.
			This helps reduce duplication
			while also leaving both functions in place.
		*/
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		/*
			In the map,
			associate the retrieved message
			with the name.
		*/
		messages[name] = message
	}
	return messages, nil
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
	// rand.Seed(time.Now().UnixNano())

	// Create a new random number generator with a custom seed (e.g., current time)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return formats[rng.Intn(len(formats))]
}
