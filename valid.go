package util

import (
	"regexp"
)

var (
	mobileRE = regexp.MustCompile(`^1(3[0-9]|4[57]|5[012356789]|7[016]|8[0-9])\d{8}$`)
	digitsRE = regexp.MustCompile(`^\d+$`)
)

func IsMobile(s string) bool {
	return mobileRE.MatchString(s)
}

func IsDigits(s string) bool {
	return digitsRE.MatchString(s)
}
