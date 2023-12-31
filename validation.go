package main

import (
	"fmt"
	"strings"
)

func sanitizeConnectCode(userId string) (string, error) {
	if !connectCodeIsValid(userId) {
		return "", fmt.Errorf("the provided userId is not valid: %s", userId)
	}

	return strings.ToUpper(userId), nil
}

func connectCodeIsValid(s string) bool {
	if s == "" {
		return false
	}

	if strings.Count(s, "#") != 1 {
		return false
	}

	return true
}
