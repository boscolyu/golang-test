package statement

import (
	"fmt"

	"github.com/bradfitz/iter"
)

var full string = ""

func main() {
	fmt.Println("Hello, 世界")
	for arg := range iter.N(10) {
		full += string(arg) + " "
		fmt.Println(full)
	}
}
