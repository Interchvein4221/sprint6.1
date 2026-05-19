package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var ErrEmptyInput = errors.New("empty input")
var ErrUnknownMorseCode = errors.New("unknown morse code")
var ErrUnknownText = errors.New("unknown text symbol")

func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", ErrEmptyInput
	}

	if isMorse(input) {
		if err := validateMorse(input); err != nil {
			return "", err
		}
		return morse.ToText(input), nil
	}

	if err := validateText(input); err != nil {
		return "", err
	}

	return morse.ToMorse(input), nil
}

func isMorse(input string) bool {
	return strings.Trim(input, ".- \n\r\t") == ""
}

func validateMorse(input string) error {
	for _, token := range strings.Fields(input) {
		if morse.MorseToRune(token) == 0 {
			return ErrUnknownMorseCode
		}
	}
	return nil
}

func validateText(input string) error {
	for _, r := range input {
		if strings.TrimSpace(string(r)) == "" {
			continue
		}
		if morse.RuneToMorse(r) == "" {
			return ErrUnknownText
		}
	}
	return nil
}
