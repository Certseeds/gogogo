package greetings

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)
import "errors"

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("input can not be empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
	_, file, _, _ := runtime.Caller(0)
	log.Printf("init %v finish \n", file)
}

func randomFormat() string {
	formats := []string{
		"Hello? %v Mi Fans",
		"Do You Like Mi For %v?",
		"Do You Like %v Band?",
		"%v thankyou very much",
	}
	return formats[rand.Intn(len(formats))]
}
