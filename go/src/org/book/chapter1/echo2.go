package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	testForRangeLoopWithIndex(os.Args)
}
var count int = 0

// count := 0
// # command-line-arguments
// ./echo2.go:21:1: syntax error: non-declaration statement outside function body

func testForRangeLoopWithIndex(list []string) {
	s, sep := "", ""

	for index, arg := range list[1:] {
		//fmt.Println(reflect.TypeOf(index))
		point := strconv.Itoa(index)
		s += sep + point + " " + arg
		sep = " "
	}
	fmt.Println(s)
}

