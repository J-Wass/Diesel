package utils

import (
	"encoding/base64"
	"os/exec"
	"runtime"
	"strings"
	"sync"
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

var (
	commitAgeOnce  sync.Once
	commitAgeValue string
)

// CommitAge returns how many commits old the current build is. The git
// lookup runs at most once per process (cached via sync.Once) since this is
// called on every /healthcheck request; shelling out on every call left
// /healthcheck vulnerable to multi-second stalls whenever the host was under
// load, which could stall any caller blocking on it.
func CommitAge() string {
	commitAgeOnce.Do(func() {
		commitAgeValue = computeCommitAge()
	})
	return commitAgeValue
}

func computeCommitAge() string {
	var out []byte
	var err error
	if runtime.GOOS == "windows" {
		out, err = exec.Command("cmd", "/C", "git", "rev-list", "--count", "master").CombinedOutput()
	} else {
		out, err = exec.Command("git", "rev-list", "--count", "master").CombinedOutput()
	}
	if err != nil {
		return "today"
	}
	return strings.TrimSpace(string(out))
}
