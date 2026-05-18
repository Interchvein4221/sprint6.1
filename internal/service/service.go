package service

import (

	"errors"
	"strings"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"

)
func isMorse(s string) bool {

	return strings.Contains(s, ".") || strings.Contains(s, "-")

}
func Convert(input string) (string, error) {

	data := strings.TrimSpace(input)
	if data == "" {
		return "", errors.New("empty input")
	}
	if isMorse(data) {
		result := morse.ToText(data)
		if result == "" {
			return "", errors.New("invalid morse")
		}
		return result, nil
	}
	result := morse.ToMorse(data)
	if result == "" {
		return "", errors.New("not convert text")
	}
	return result, nil

}