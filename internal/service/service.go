package service

import (

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) string {

	if len(data) == 0 {
		return ""
	}
	if containsMorse(data) {
		return morse.ToText(data)
	}
	return morse.ToMorse(data)

}

func containsMorse(s string) bool {
	for _, r := range s {
		if r == '.' || r == '-' {
			return true
		}
	}
	return false
}