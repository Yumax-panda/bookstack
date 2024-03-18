package random

import (
	crand "crypto/rand"
	"io"
)

// Salt 64bytesのランダムなソルトを生成する
func Salt() []byte {
	salt := make([]byte, 64)
	_, _ = io.ReadFull(crand.Reader, salt)
	return salt
}
