package service

import (
	"strings"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {

	data = strings.TrimSpace(data)
	if data == "" {
		return "", nil
	}
	if strings.Contains(data, ".") || strings.Contains(data, "-") {
		return morse.ToText(data), nil
	}
	return morse.ToMorse(data), nil

}
