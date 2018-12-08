package statement
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

	if a == b {
		fmt.Println(EQUAL)
	} else {
		fmt.Println(NOT_EQUAL)
	}

	if a == c {
		fmt.Println(EQUAL)
	} else {
		fmt.Println(NOT_EQUAL)
	}
		
	
}
