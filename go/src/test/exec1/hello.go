package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	println("What is your name?")
	var data, error = getName()
	if error != nil {
		println(error)
		os.Exit(-1);
	}
	printName(makeString(data))
}

func getName() (string, error) {
	return bufio.NewReader(os.Stdin).ReadString('\n')
}

func makeString(name string) (string) {
	var splitData = strings.Split(name, "\n")
	return "Hello, " + splitData[0] + ", nice to meet you!"
}

func printName(message string) {
	println(message)
}
