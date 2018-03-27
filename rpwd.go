package rpwd

import (
	"math/rand"
	"time"
)

func Init(sLength int, sUpper bool, sLower bool, sNumber bool, sCharacter bool) []byte {
	var str = ""

	if sUpper {
		str += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if sLower {
		str += "abcdefghijklmnopqrstuvwxyz"
	}

	if sNumber {
		str += "0123456789"
	}

	if sCharacter {
		str += "!@#$%^&*()"
	}

	if !sUpper && !sLower && !sNumber && !sCharacter {
		str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < sLength; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}
