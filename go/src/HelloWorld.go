// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"com"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Hello, 世界")
	variable()
	com.PackageTest()
}

func variable() {
	// var 로 시작함.
	// 변수명 뒤에 자료형이 붙음
	// long이라는 자료형은 없음.
	var intVar int = 100
	var longVar float32 = 100
	// java 처럼 String 이라는 자료형이 있음
	//var stringVar string = "test string"

	fmt.Println(intVar)
	fmt.Println(longVar)
	//fmt.Println(stringVar + " " + string(longVar))
}

func condition() {
	var conditionVar int = 1
	if conditionVar == 1 {
		fmt.Println("test variable")
	}
	/*else if conditionVar ==2 {
		fmt.Println("test variable2")
	}

	if 100 == 100 {
		fmt.Println("100 is equal")
	}*/

}


func test() {
	fmt.Println("Hello, 世界!")

	for argkey, arg := range os.Args {
		fmt.Print(argkey)
		fmt.Println(" " + arg)
	}

	counts := make(map[string]int)
	fmt.Println("run the program\n")
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	fmt.Printf("map data size %v\n", len(counts))

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}