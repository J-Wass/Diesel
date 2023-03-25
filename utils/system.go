package utils

import (
	"encoding/base64"
	"os/exec"
	"runtime"
)

// Decodes base64 strings.
func DecodedFromBase64(base64String string) (string, error) {
	decodedUrl, err := base64.StdEncoding.DecodeString(base64String)
	return string(decodedUrl), err
}

// Encodes base64 strings
func EncodedBase64(inputString string) string {
	return base64.StdEncoding.EncodeToString([]byte(inputString))
}

func CommitAge() string{
	if runtime.GOOS == "windows" {
		commitAge, err := exec.Command("cmd", "/C","git", "rev-list", "--count master").CombinedOutput()
		if err != nil{
			return "today"
		}
		return string(commitAge)
    } else {
        commitAge, err := exec.Command("git", "rev-list", "--count master").CombinedOutput()
		if err != nil{
			return "today"
		}
		return string(commitAge)
    }
}