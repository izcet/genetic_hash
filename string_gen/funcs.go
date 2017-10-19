package main 

import (
	"strings"
	"math/rand"
	"bytes"
	"strconv"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func	rand40() string {
	var buffer bytes.Buffer

	for i := 0; i < 40; i++ {
		buffer.WriteString(string(letters[rand.Intn(len(letters))]))
	}
	return (buffer.String())
}

var lotsA = repeatA()

func	randA() string {
	str := lotsA("a")
	return (str)
}

func	repeatA() func(string) string {
	strLen := 0
	return func(str string) string {
		strLen += 1
		return (strings.Repeat(str, strLen))
	}
}

func	randRandAB() string {
	var buffer bytes.Buffer
	len := rand.Intn(80)
	for i := 0; i <= len; i++ {
		buffer.WriteString(string(letters[rand.Intn(2)]))
	}
	return (buffer.String())
}

func	fooBarBaz() string {
	var buffer bytes.Buffer
	len := rand.Intn(20)
	for i := 0; i <= len; i++ {
		x := rand.Intn(4)
		var str string
		switch x {
		case 0: 
			str = "foo"
		case 1:
			str = "bar"
		case 2:
			str = "baz"
		case 3:
			str = " "
		}
		buffer.WriteString(str)
	}
	return (buffer.String())
}

var lotsFizz = repeatFizz()

func	fizzBuzz() string {
	str := lotsFizz()
	return (str)
}

func	repeatFizz() func() string {
	strLen := 0
	var buffer bytes.Buffer
	return func() string {
		strLen += 1
		if (strLen != 1) {
			buffer.WriteString(", ")
		}
		fb := 0
		if (strLen % 3) == 0 {
			buffer.WriteString("fizz")
			fb += 1
		}
		if (strLen % 5) == 0 {
			buffer.WriteString("buzz")
			fb += 1
		}
		if (fb == 0) {
			buffer.WriteString(strconv.Itoa(strLen))
		}
		result := buffer.String() + "\n"
		return (result)
	}
}

