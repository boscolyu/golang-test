package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""
	fmt.Println(time.Now())
	count := strconv.Itoa(os.Args[1])

	for {
		s = sep + count
		if count > 100 {
			break
		}
	}
	fmt.Println(s)
	fmt.Println(time.Now())


	fmt.Println(strings.Join(os.Args[2:], " "))
	fmt.Println(time.Now())
	fmt.Println(os.Args)
	fmt.Println(time.Now())
}