package main

import (
	"fmt"
	"os"
)

func	setFile(filename string) (*os.File, error) {
	err := os.Remove(filename)
	if (err == nil) {
		fmt.Println("Deleted file:", filename)
	}
	return (os.Create(filename))
}

func	main() {
	global := "global_list.txt"
	numStrings := 256
	
	file, err := setFile(global)
	if (err != nil) {
		fmt.Println("Global List File Error:", err)
		return //()
	}
	defer file.Close()
	fmt.Println("Created", global)

	makeFileList("40 random characters", "rand40.txt", rand40, file, numStrings)
	makeFileList("increasing numbers of 'a'", "randA.txt", randA, file, numStrings)
	makeFileList("random combinations of 'a' and 'b' of random lengths", "randRandAB.txt", randRandAB, file, numStrings)
	makeFileList("foos, bars, and bazs", "fooBarBaz.txt", fooBarBaz, file, numStrings)
	makeFileList("fizzbuzz iterations", "fizzBuzz.txt", fizzBuzz, file, numStrings)
}

func	makeFileList(description, filename string, function func() string, global *os.File, numStrings int) {
	err := singleFunction(description, filename, function, numStrings)
	if (err != nil) {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Made list of", description)
		err := writeLine(filename, global)
		if (err != nil) {
			panic(err)
		}
	}
}

func	singleFunction(description, filename string, function func() (string), numStrings int) error {
	file, err := setFile(filename)
	if (err != nil)	{
		return (err)
	}
	defer file.Close()
	description = "# " + description
	err = writeLine(description, file)
	if (err != nil) {
		return (err)
	}
	stringMap := getList(function, numStrings)
	err = writeListToFile(stringMap, file)
	return (err)
}

func	getList(function func() string, numStrings int) map[string]int {
	list := make(map[string]int)
	for ; len(list) < numStrings; {
		str := function()
		//fmt.Println(str) // debugging individual strings
		if _, ok := list[str]; ok == false {
			list[str] = 0
		}
	}
	return (list)
}

func	writeListToFile(list map[string]int, file *os.File) error {
	for str, _ := range list {
		err := writeLine(str, file)
		if (err != nil) {
			return (err)
		}
	}
	return (nil)
}

func	writeLine(str string, file *os.File) error {
	_, err := file.WriteString(str)
	if (err != nil) {
		return (err)
	}
	_, err = file.WriteString("\n")
	if (err != nil) {
		return (err)
	}
	return (nil)
}
