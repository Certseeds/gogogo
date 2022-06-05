package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello? %s Thank you! Thank you very much", name)
	return message
}
