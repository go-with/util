package util

import (
	"regexp"
)

var (
	mobileRE = regexp.MustCompile(`^1(3[0-9]|4[1456789]|5[012356789]|6[6]|7[01345678]|8[0-9]|9[189])\d{8}$`)
	digitsRE = regexp.MustCompile(`^\d+$`)
)

func IsMobile(s string) bool {
	return mobileRE.MatchString(s)
}

func IsDigits(s string) bool {
	return digitsRE.MatchString(s)
}
