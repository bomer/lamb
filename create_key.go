package main

import (
	"crypto/sha256"
	"fmt"
)

func CreateKey(name string) string {
	sum := sha256.Sum256([]byte(name))
	return fmt.Sprintf("%x", sum)
}
