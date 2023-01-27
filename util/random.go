package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcedfghijklmnopqrstuvwxyz"

func init() {
    rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
    // Int63n will not use the negative value    
    return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string { 
    var stringBuilder strings.Builder
    lengthOfAlphabet := len(alphabet)

    for i := 0; i < n; i++ {
        c := alphabet[rand.Intn(lengthOfAlphabet)]
        stringBuilder.WriteByte(c)
    }
    return stringBuilder.String()
}

func RandomOwner() string {
    return RandomString(6)
}
