package utils

import "encoding/base64"

// Decodes base64 strings.
func DecodedFromBase64(base64String string) (string, error) {
	decodedUrl, err := base64.StdEncoding.DecodeString(base64String)
	return string(decodedUrl), err
}

// Encodes base64 strings
func EncodedBase64(inputString string) string {
	return base64.StdEncoding.EncodeToString([]byte(inputString))
}