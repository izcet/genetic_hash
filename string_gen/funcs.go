package main 

import (
	"strings"
	"math/rand"
	"bytes"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func	rand40() string {
	var buffer bytes.Buffer

	for i := 0; i < 40; i++ {
		buffer.WriteString(string(letters[rand.Intn(len(letters))]))
	}
	return (buffer.String())
}

var fun = repeat()

func	randA() string {
	str := fun("a")
	return (str)
}

func	repeat() func(string) string {
	strLen := 0
	return func(str string) string {
		strLen += 1
		return (strings.Repeat(str, strLen))
	}
}
