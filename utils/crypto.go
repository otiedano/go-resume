package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

//EncodeSHA256 sha256 encryption
func EncodeSHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//EncodeSHA256ByByte sha256 encryption
func EncodeSHA256ByByte(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

//GenToken genarate token
func GenToken(s string) string {
	time := time.Now().String()
	sum := time + s
	return EncodeSHA256(sum)
}
