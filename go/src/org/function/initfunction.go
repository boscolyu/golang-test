package main

import "fmt"

var (
	// function is not running just to put in the variable.
	funcVar = run
)
func main() {
	fmt.Println("1. main function")
	run()
	funcVar()
}

// init function is runned before main function.
func init() {
	fmt.Println("0. test init function")
}


func run() {
	fmt.Println("2. test run fuction")
}