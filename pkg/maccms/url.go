package maccms

import (
	"fmt"
	"strings"
)

func EncodeURL(rawURL string) string {
	var encodedURL strings.Builder
	for _, char := range rawURL {
		if isAlphaNumeric(char) || char == '.' || char == '/' || char == ':' || char == '_' {
			encodedURL.WriteString(fmt.Sprintf("%%%02X", char))
		} else {
			encodedURL.WriteRune(char)
		}
	}
	return encodedURL.String()
}

func isAlphaNumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}
