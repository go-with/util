package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func MD5Str(s string) string {
	cs := md5.Sum([]byte(s))
	return hex.EncodeToString(cs[:])
}

func SHA1Str(s string) string {
	cs := sha1.Sum([]byte(s))
	return hex.EncodeToString(cs[:])
}
