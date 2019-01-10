package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// for (int i = 1; i < len(os.Args); i++) { }
	// c, java style loop statement
	// loop statement of golang is different from c, java

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	testForStatementWithoutCondition()

	testOperator()
}

func testOperator() {

	// It is possible to define and init variables with := keyword.
	k, j := 1, 3
	j = j + 1
	j++
	// j = j ++
	// ./echo1.go:20:8: syntax error: unexpected ++ at end of statement

	fmt.Println(j, k)

}
func testForStatementWithoutCondition() {

	count := 0
	for {
		fmt.Println( "statement : " , count)
		if count > 100 {
			break;
		}
		count ++
	}
}
