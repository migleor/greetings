package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hola, %v! Bienvenido!",
		"!Que bueno verte, %v!",
		"¡Saludos amigo, %v! ¿Cómo estás?",
		"¡Hola, %v! ¿Qué tal?",
		"¡Hola, %v! ¿Cómo te va?",
	}

	return formats[rand.Intn(len(formats))]
}
