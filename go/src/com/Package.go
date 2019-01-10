package main

import (

	"com/bosco/package1"
	"com/bosco/package2"
	"fmt"
)

func main() {
	fmt.Println("package test")
	fmt.Println(package1.TestPackage1())
	fmt.Println(package2.TestPackage2())

}
