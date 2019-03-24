package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("What is the input string?")
	var data, error = getString()
	if error != nil {
		println(error)
		os.Exit(-1);
	}
	printString(makeString(data))
}

func getString() (string, error) {
	return bufio.NewReader(os.Stdin).ReadString('\n')
}

func makeString(name string) (string) {
	var splitData = strings.Split(name, "\n")
	println(splitData[0] + " has " + string(len(splitData[0])) + " characters.")
	return splitData[0] + " has " + strconv.Itoa(len(splitData[0])) + " characters."
}

func printString(message string) {
	println(message)
}