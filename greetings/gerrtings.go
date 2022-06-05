package greetings

import "fmt"
import "errors"

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Input can not be empty")
	}
	message := fmt.Sprintf("Hello? %s Thank you! Thank you very much", name)
	return message, nil
}
