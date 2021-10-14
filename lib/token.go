package lib

import (
	"crypto/rand"
	"fmt"
)

func CreateToken() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)[:16]
}
