package main
import (
	"fmt"
)

const (
        EQUAL = iota
        NOT_EQUAL
)

func main() {
	var a = "string a"
	var b = "string b"
	var c = "string a"
	
	var map1 = make(map[string]int)
	var map2 = make(map[string]int)

	if map1 == map1 {
		fmt.Println(EQUAL)
	} else {
		fmt.Println(NOT_EQUAL)
	}

	if map1 == map2 {
		fmt.Println(EQUAL)
	} else {
		fmt.Println(NOT_EQUAL)
	}

	map1[a] = 1
	map2[a] = 1

	if map1 == map1 {
		fmt.Println(EQUAL)
	} else {
		fmt.Println(NOT_EQUAL)
	}

	if map1 == map2 {
		fmt.Println(EQUAL)
	} else {
		fmt.Println(NOT_EQUAL)
	}


}
