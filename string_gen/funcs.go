package main 

import (
	"strings"
	"math/rand"
	"bytes"
	"strconv"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const alnumer = letters + numbers

var firstNames = [...]string{
	"alice", "alex", "bob", "betty", "chris", "carol", "dave", "daniela", "edward", "emily", "frank", "florence",
	"gabby", "gabe", "heather", "hubert", "isabella", "isaac", "james", "janice", "konstantine", "kelly",
	"leonard", "lily", "meghan", "mike", "nathan", "natalie", "oscar", "olivia", "patricia", "paul", "quentin", "queenie",
	"romeo", "rachel", "samantha", "sebastian", "terri", "timothy", "ulysses", "ursula", "victor", "victoria",
	"winston", "wendy", "xavier", "xena", "yolanda", "yeezy", "zachary", "zoey" }
var domains = [...]string{
	"google", "apple", "yahoo", "hotmail", "sharklasers", "website", "example", "gmail", "facebook" }
var extensions = [...]string{
	"net", "com", "us", "org", "mail", "uk" }

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

func	randAB() string {
	var buffer bytes.Buffer
	strLen := rand.Intn(80)
	for i := 0; i <= strLen; i++ {
		buffer.WriteString(string(letters[rand.Intn(2)]))
	}
	return (buffer.String())
}

func	fooBarBaz() string {
	var buffer bytes.Buffer
	strLen := rand.Intn(20)
	for i := 0; i <= strLen; i++ {
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
		return (buffer.String())
	}
}

func	herpDerp() string {
	var buffer bytes.Buffer
	prefixes := []string{"b", "bl", "d", "f", "fl", "h", "h", "l", "m", "n", "p", "s", "sh", "sl", "t"}
	middles := []string{"erp", "ee", "erp", "er", "erp", "oo"}
	suffixes := []string{"ler", "s", "", "us", "um", "er", ""}
	strLen:= rand.Intn(10)
	for i := 0; i <= strLen; i++ {
		if (i != 0) {
			buffer.WriteString(" ")
		}
		x := rand.Intn(len(prefixes))
		buffer.WriteString(prefixes[x])
		x = rand.Intn(len(middles))
		buffer.WriteString(middles[x])
		x = rand.Intn(len(suffixes))
		buffer.WriteString(suffixes[x])
	}
	return (buffer.String())
}

func	alphaNum() string {
	var buffer bytes.Buffer
	strLen := rand.Intn(80)
	for i := 0; i <= strLen; i++ {
		buffer.WriteString(string(alnumer[rand.Intn(len(alnumer))]))
	}
	return (buffer.String())
}

func	digits() string {
	var buffer bytes.Buffer
	strLen := rand.Intn(80)
	for i := 0; i <= strLen; i++ {
		buffer.WriteString(string(numbers[rand.Intn(len(numbers))]))
	}
	return (buffer.String())
}

var lotsNums = repeatNums()

func	counters() string {
	str := lotsNums()
	return (str)
}

func	repeatNums() func () string {
	strLen := 0
	var buffer bytes.Buffer
	return func() string {
		strLen += 1
		buffer.WriteString(strconv.Itoa(strLen))
		return (buffer.String())
	}
}

func	getRandomFromArr(arr []int) int {
	return (arr[rand.Intn(len(arr))])
}

func	emails() string {
	var buffer bytes.Buffer
	var arr = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 3}
	numNames := getRandomFromArr(arr)
	for x := 0; x < numNames; x++ {
		buffer.WriteString(firstNames[rand.Intn(len(firstNames))])
	}
	numNums := rand.Intn(5)
	for x := 0; x < numNums; x++ {
		buffer.WriteString(string(numbers[rand.Intn(len(numbers))]))
	}
	buffer.WriteString("@")
	buffer.WriteString(domains[rand.Intn(len(domains))])
	numEnds := getRandomFromArr(arr)
	for x := 0; x < numEnds; x++ {
		buffer.WriteString(".")
		whichEnd := rand.Intn(len(extensions) + 1)
		if (whichEnd == len(extensions)) {
			for x := 0; x < 2; x++ {
				buffer.WriteString(string(letters[rand.Intn(26)]))
			}
		} else {
			buffer.WriteString(extensions[whichEnd])
		}
	}
	return (buffer.String())
}
