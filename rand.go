package util

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	alnum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var (
	stdRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func RandNumStr(n int) string {
	return string(RandNum(n))
}

func RandNum(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = strconv.Itoa(stdRand.Intn(10))[0]
	}
	return b
}

func RandAlnumStr(n int) string {
	return string(RandAlnum(n))
}

func RandAlnum(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = alnum[stdRand.Intn(len(alnum))]
	}
	return b
}
