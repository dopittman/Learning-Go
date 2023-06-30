package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("Hello: empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomHelloFormat(), name)
	return message, nil
}

func Goodbye(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("Goodbye: empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomGoodbyeFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomHelloFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomGoodbyeFormat() string {
	// A slice of message formats.
	formats := []string{
		"Later, %v.",
		"Adios, %v!",
		"Hasta la vista, %v!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
