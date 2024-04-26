package helpers

import (
	"math/rand"
	"time"
)

func GenerateVA() string {
	rand.Seed(time.Now().UnixNano())

	vaLength := 16

	vaDigits := make([]byte, vaLength)

	for i := 0; i < vaLength; i++ {
		vaDigits[i] = byte(rand.Intn(10)) + '0'
	}

	return string(vaDigits)
}
