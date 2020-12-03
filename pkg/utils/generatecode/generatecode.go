package generatecode

import (
	"crypto/rand"
	"io"
	"strconv"
)

func ActivationCode(max int) int32 {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	var data, _ = strconv.Atoi(string(b))
	return int32(data)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
