/*
https://go.dev/doc/tutorial/add-a-test
Ending a file's name with _test.go
tells the go test command
that this file contains test functions
*/
package greetings

import (
	"math/rand"
	"regexp"
	"testing"
	"time"
)

/*
"TestHelloName"

	calls
		"greetings.Hello"
		with a name, checking for a valid return value.
*/
func TestHelloName(t *testing.T) {

	// Create a new random number generator with a custom seed (e.g., current time)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	name := "Mater"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name, rng)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(
			"Hello(%q) = %q, %v, want match for %#q, nil",
			msg,
			msg,
			err,
			want,
		)
	}
}

/*
TestHelloEmpty calls "greetings.Hello" with an empty string,
checking for an error.
*/
func TestHelloEmpty(t *testing.T) {
	// Create a new random number generator with a custom seed (e.g., current time)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	name := ""
	_, err := Hello(name, rng)
	if err != nil {
		t.Fatalf(
			`An error occured : %q. Provide a real first name, please.`,
			err,
		)
	}
}

/*
Add a test

Now that you've got your code to a stable place, add a test.
Testing your code during development may expose bugs
that find their way in as you make changes.
In this topic,
you add a test for the Hello function:

func Hello(name string, rng *rand.Rand) (string, error)

Go's built-in support for unit tests
makes it easier to test as you go.
Specifically, using naming conventions,
Go's
- "testing package", and the go
- "test command",
you can quickly write and execute tests.

1.
In the greetings directory, create a file called "greetings_test.go".
Ending a file's name with _test.go tells the go test command
that this file contains test functions.

2.
In greetings_test.go, paste the following code and save the file:
above

https://pkg.go.dev/testing/#T.Fatalf

 In this code, you:

    - Implement test functions in the same package as the code you're testing.
    - Create two test functions to test the "greetings.Hello" function.
	  Test function names have the form TestName,
	  Name says something about the specific test.
	  Also, test functions take a pointer
	  to the testing package's testing.T type as a parameter.
	  You use this parameter's methods for reporting and logging from your test.

    - Implement two tests:
        "TestHelloName" calls the Hello function,
		passing a name value
		with which the function should be able
		to return a valid response message.
		If the call returns an error or an unexpected response message
		(one that doesn't include the name you passed in),
		you use "the t parameter's Fatalf" method
		to print a message to the console and end execution.

		"TestHelloEmpty" calls the Hello function with an empty string.
		This test is designed to confirm
		that your error handling works.
		If the call returns a non-empty string or no error,
		you use the t parameter's Fatalf method
		to print a message to the console and end execution.

3.
At the command line in the greetings directory,
run the "go test" command to execute the test.
https://go.dev/cmd/go/#hdr-Test_packages


The "go test" command executes test "functions", whose names begin with "Test",
e.g. "TestHelloEmpty", "TestHelloName",
in test "files", whose names end with _test.go,
e.g. "greetins_test.go".

You can add the -v flag to get verbose output
that lists all of the tests and their results.

The tests should pass.


*/
