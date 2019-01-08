package main

import (
	"com"
	package5 "com/bosco/package1"
	"com/bosco/package2"
	"com/bosco/package3"
	"fmt"
)

func BoscoFunction() {
	fmt.Println("org.package.BoscoFunction")
	bosco.BoscoFunction()
	package3.TestPackage3()
	package4.TestPackage2()
	package5.TestPackage1()
}

func main() {
	BoscoFunction()
}