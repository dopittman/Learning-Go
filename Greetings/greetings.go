package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("Hello: empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func Goodbye(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("Goodbye: empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Goodbye, %v!", name)
	return message, nil
}
