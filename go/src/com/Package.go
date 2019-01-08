package bosco

import (
	// just add the folder
	package5 "com/bosco/package1"
	"com/bosco/package2"
	"com/bosco/package3"
	"fmt"
)

func main() {
	fmt.Println("package test")
	// referenced the package name
	fmt.Println(package3.TestPackage3())
	fmt.Println(package4.TestPackage2())
	fmt.Println(package5.TestPackage1())

}

func BoscoFunction() bool {
	main()
	if package5.TestPackage1() == package3.TestPackage3() {
		return false
	}
	return true
}
