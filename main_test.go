package main_test

import (
	"fmt"
	"goapi/pkg/maccms"
	"net/url"
	"testing"
)

func TestBase64(t *testing.T) {
	ASAS := "%68%74%74%70%73%3A%2F%2F%76%2E%63%64%6E%6C%7A%32%2E%63%6F%6D%2F%32%30%32%34%30%32%30%35%2F%32%38%37%30%38%5F%64%35%34%66%37%66%61%64%2F%69%6E%64%65%78%2E%6D%33%75%38"
	encoded := maccms.Base64encode(ASAS)
	fmt.Println("encoded", encoded) // Output: SGVsbG8sIOS4lueVjA==

	decoded := maccms.Base64decode(encoded)
	fmt.Println("decoded", decoded) // Output: Hello, 世界!
}

func TestUrl(t *testing.T) {
	Urls()
}

func Urls() {
	encodedURL := "%68%74%74%70%73%3A%2F%2F%76%2E%63%64%6E%6C%7A%32%2E%63%6F%6D%2F%32%30%32%34%30%32%30%35%2F%32%38%37%30%38%5F%64%35%34%66%37%66%61%64%2F%69%6E%64%65%78%2E%6D%33%75%38"

	decodedURL, err := url.QueryUnescape(encodedURL)
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return
	}

	fmt.Println("Decoded URL:", decodedURL)
	encodedURL2 := maccms.EncodeURL(decodedURL)

	fmt.Println("EncodedURL URL:", encodedURL2)
}
