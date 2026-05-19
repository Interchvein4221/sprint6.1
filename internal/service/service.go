package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var ErrEmptyInput = errors.New("Ошибка ввода")
var ErrUnknownMorseCode = errors.New("Неизвестный код азбуки Морзе")
var ErrUnknownText = errors.New("Неизвестный буква алфавита")

func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", ErrEmptyInput
	}

	if strings.Trim(input, ".- \n\r\t") == "" {
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
