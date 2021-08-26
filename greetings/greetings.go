package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.

    if name == "" {
        return "", errors.New("empty name")
    }


    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}